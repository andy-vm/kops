/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsmodel

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/dns"
	"k8s.io/kops/pkg/model"
	"k8s.io/kops/pkg/model/iam"
	"k8s.io/kops/pkg/util/stringorslice"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/awstasks"
)

// IAMModelBuilder configures IAM objects
type IAMModelBuilder struct {
	*AWSModelContext
	Lifecycle *fi.Lifecycle
	Cluster   *kops.Cluster
}

var _ fi.ModelBuilder = &IAMModelBuilder{}

const NodeRolePolicyTemplate = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": { "Service": "{{ IAMServiceEC2 }}"},
      "Action": "sts:AssumeRole"
    }
  ]
}`

func (b *IAMModelBuilder) Build(c *fi.ModelBuilderContext) error {
	// Collect managed Instance Group roles
	managedRoles := make(map[kops.InstanceGroupRole]bool)

	// Collect Instance Profile ARNs and their associated Instance Group roles
	sharedProfileARNsToIGRole := make(map[string]kops.InstanceGroupRole)
	for _, ig := range b.InstanceGroups {
		if ig.Spec.IAM != nil && ig.Spec.IAM.Profile != nil {
			specProfile := fi.StringValue(ig.Spec.IAM.Profile)
			if matchingRole, ok := sharedProfileARNsToIGRole[specProfile]; ok {
				if matchingRole != ig.Spec.Role {
					return fmt.Errorf("found IAM instance profile assigned to multiple Instance Group roles %v and %v: %v",
						ig.Spec.Role, sharedProfileARNsToIGRole[specProfile], specProfile)
				}
			} else {
				sharedProfileARNsToIGRole[specProfile] = ig.Spec.Role
			}
		} else {
			managedRoles[ig.Spec.Role] = true
		}
	}

	// Generate IAM tasks for each shared role
	for profileARN, igRole := range sharedProfileARNsToIGRole {
		lchPermissions := false
		defaultWarmPool := b.Cluster.Spec.WarmPool
		for _, ig := range b.InstanceGroups {
			warmPool := defaultWarmPool.ResolveDefaults(ig)
			if ig.Spec.Role == igRole && warmPool.IsEnabled() && warmPool.EnableLifecycleHook {
				lchPermissions = true
				break

			}
		}
		role, err := iam.BuildNodeRoleSubject(igRole, lchPermissions)
		if err != nil {
			return err
		}

		iamName, err := model.FindCustomAuthNameFromArn(profileARN)
		if err != nil {
			return fmt.Errorf("unable to parse instance profile name from arn %q: %v", profileARN, err)
		}
		err = b.buildIAMTasks(role, iamName, c, true)
		if err != nil {
			return err
		}
	}

	// Generate IAM tasks for each managed role
	defaultWarmPool := b.Cluster.Spec.WarmPool
	for igRole := range managedRoles {
		haveWarmPool := false
		for _, ig := range b.InstanceGroups {
			warmPool := defaultWarmPool.ResolveDefaults(ig)
			if ig.Spec.Role == igRole && warmPool.IsEnabled() && warmPool.EnableLifecycleHook {
				haveWarmPool = true
				break

			}
		}
		role, err := iam.BuildNodeRoleSubject(igRole, haveWarmPool)
		if err != nil {
			return err
		}

		iamName := b.IAMName(igRole)
		if err := b.buildIAMTasks(role, iamName, c, false); err != nil {
			return err
		}
	}

	iamSpec := b.Cluster.Spec.IAM
	if iamSpec != nil {
		for _, sa := range iamSpec.ServiceAccountExternalPermissions {
			var p *iam.Policy
			aws := sa.AWS
			if aws.InlinePolicy != "" {
				bp, err := b.buildPolicy(aws.InlinePolicy)
				p = bp
				if err != nil {
					return fmt.Errorf("error inline policy: %w", err)
				}
			}
			serviceAccount := &iam.GenericServiceAccount{
				NamespacedName: types.NamespacedName{
					Name:      sa.Name,
					Namespace: sa.Namespace,
				},
				Policy: p,
			}
			iamRole, err := b.BuildServiceAccountRoleTasks(serviceAccount, c)
			if err != nil {
				return fmt.Errorf("error building service account role tasks: %w", err)
			}
			if len(aws.PolicyARNs) > 0 {
				name := "external-" + fi.StringValue(iamRole.Name)
				externalPolicies := aws.PolicyARNs
				c.AddTask(&awstasks.IAMRolePolicy{
					Name:             fi.String(name),
					ExternalPolicies: &externalPolicies,
					Managed:          true,
					Role:             iamRole,
					Lifecycle:        b.Lifecycle,
				})
			}
		}

	}

	return nil
}

// BuildServiceAccountRoleTasks build tasks specifically for the ServiceAccount role.
func (b *IAMModelBuilder) BuildServiceAccountRoleTasks(role iam.Subject, c *fi.ModelBuilderContext) (*awstasks.IAMRole, error) {
	iamName, err := b.IAMNameForServiceAccountRole(role)
	if err != nil {
		return nil, err
	}

	iamRole, err := b.buildIAMRole(role, iamName, c)
	if err != nil {
		return nil, err
	}

	if err := b.buildIAMRolePolicy(role, iamName, iamRole, c); err != nil {
		return nil, err
	}

	return iamRole, nil
}

func (b *IAMModelBuilder) buildIAMRole(role iam.Subject, iamName string, c *fi.ModelBuilderContext) (*awstasks.IAMRole, error) {
	roleKey, isServiceAccount := b.roleKey(role)

	rolePolicy, err := b.buildAWSIAMRolePolicy(role)
	if err != nil {
		return nil, err
	}

	iamRole := &awstasks.IAMRole{
		Name:      fi.String(iamName),
		Lifecycle: b.Lifecycle,

		RolePolicyDocument: rolePolicy,
		Tags:               b.CloudTags(iamName, false),
	}

	if isServiceAccount {
		// e.g. kube-system-dns-controller
		iamRole.ExportWithID = fi.String(roleKey)
	} else {
		// e.g. nodes
		iamRole.ExportWithID = fi.String(roleKey + "s")
	}

	if b.Cluster.Spec.IAM != nil && b.Cluster.Spec.IAM.PermissionsBoundary != nil {
		iamRole.PermissionsBoundary = b.Cluster.Spec.IAM.PermissionsBoundary
	}

	c.AddTask(iamRole)

	return iamRole, nil
}

func (b *IAMModelBuilder) buildIAMRolePolicy(role iam.Subject, iamName string, iamRole *awstasks.IAMRole, c *fi.ModelBuilderContext) error {
	iamPolicy := &iam.PolicyResource{
		Builder: &iam.PolicyBuilder{
			Cluster:              b.Cluster,
			Role:                 role,
			Region:               b.Region,
			UseServiceAccountIAM: b.UseServiceAccountIAM(),
		},
	}

	if !dns.IsGossipHostname(b.Cluster.ObjectMeta.Name) {
		// This is slightly tricky; we need to know the hosted zone id,
		// but we might be creating the hosted zone dynamically.
		// We create a stub-reference which will be combined by the execution engine.
		iamPolicy.DNSZone = &awstasks.DNSZone{
			Name: fi.String(b.NameForDNSZone()),
		}
	}

	t := &awstasks.IAMRolePolicy{
		Name:      fi.String(iamName),
		Lifecycle: b.Lifecycle,

		Role:           iamRole,
		PolicyDocument: iamPolicy,
	}
	c.AddTask(t)

	return nil
}

// roleKey builds a string to represent the role uniquely.  It returns true if this is a service account role.
func (b *IAMModelBuilder) roleKey(role iam.Subject) (string, bool) {
	serviceAccount, ok := role.ServiceAccount()
	if ok {
		return strings.ToLower(serviceAccount.Namespace + "-" + serviceAccount.Name), true
	}

	// This isn't great, but we have to be backwards compatible with the old names.
	switch role.(type) {
	case *iam.NodeRoleMaster:
		return strings.ToLower(string(kops.InstanceGroupRoleMaster)), false
	case *iam.NodeRoleAPIServer:
		return strings.ToLower(string(kops.InstanceGroupRoleAPIServer)), false
	case *iam.NodeRoleNode:
		return strings.ToLower(string(kops.InstanceGroupRoleNode)), false
	case *iam.NodeRoleBastion:
		return strings.ToLower(string(kops.InstanceGroupRoleBastion)), false

	default:
		klog.Fatalf("unknown node role type: %T", role)
		return "", false
	}
}

func (b *IAMModelBuilder) buildIAMTasks(role iam.Subject, iamName string, c *fi.ModelBuilderContext, shared bool) error {
	roleKey, _ := b.roleKey(role)

	iamRole, err := b.buildIAMRole(role, iamName, c)
	if err != nil {
		return err
	}

	if err := b.buildIAMRolePolicy(role, iamName, iamRole, c); err != nil {
		return err
	}

	{
		// To minimize diff for easier code review

		var iamInstanceProfile *awstasks.IAMInstanceProfile
		{
			iamInstanceProfile = &awstasks.IAMInstanceProfile{
				Name:      fi.String(iamName),
				Lifecycle: b.Lifecycle,
				Shared:    fi.Bool(shared),
				Tags:      b.CloudTags(iamName, false),
			}
			c.AddTask(iamInstanceProfile)
		}

		{
			iamInstanceProfileRole := &awstasks.IAMInstanceProfileRole{
				Name:      fi.String(iamName),
				Lifecycle: b.Lifecycle,

				InstanceProfile: iamInstanceProfile,
				Role:            iamRole,
			}
			c.AddTask(iamInstanceProfileRole)
		}

		// Create External Policy tasks
		if !shared {
			var externalPolicies []string

			if b.Cluster.Spec.ExternalPolicies != nil {
				p := *(b.Cluster.Spec.ExternalPolicies)
				externalPolicies = append(externalPolicies, p[roleKey]...)
			}
			sort.Strings(externalPolicies)

			name := fmt.Sprintf("%s-policyoverride", roleKey)
			t := &awstasks.IAMRolePolicy{
				Name:             fi.String(name),
				Lifecycle:        b.Lifecycle,
				Role:             iamRole,
				Managed:          true,
				ExternalPolicies: &externalPolicies,
			}

			c.AddTask(t)
		}

		// Generate additional policies if needed, and attach to existing role
		if !shared {
			additionalPolicy := ""
			if b.Cluster.Spec.AdditionalPolicies != nil {
				additionalPolicies := *(b.Cluster.Spec.AdditionalPolicies)

				additionalPolicy = additionalPolicies[roleKey]
			}

			additionalPolicyName := "additional." + iamName

			t := &awstasks.IAMRolePolicy{
				Name:      fi.String(additionalPolicyName),
				Lifecycle: b.Lifecycle,

				Role: iamRole,
			}

			if additionalPolicy != "" {
				p, err := b.buildPolicy(additionalPolicy)
				if err != nil {
					return fmt.Errorf("additionalPolicy %q is invalid: %v", roleKey, err)
				}

				policy, err := p.AsJSON()
				if err != nil {
					return fmt.Errorf("error building IAM policy: %w", err)
				}

				t.PolicyDocument = fi.NewStringResource(policy)
			} else {
				t.PolicyDocument = fi.NewStringResource("")
			}

			c.AddTask(t)
		}
	}

	return nil
}

func (b *IAMModelBuilder) buildPolicy(policyString string) (*iam.Policy, error) {
	p := &iam.Policy{
		Version: iam.PolicyDefaultVersion,
	}

	statements, err := iam.ParseStatements(policyString)
	if err != nil {
		return nil, err
	}

	p.Statement = append(p.Statement, statements...)
	return p, nil
}

// IAMServiceEC2 returns the name of the IAM service for EC2 in the current region.
// It is ec2.amazonaws.com in the default aws partition, but different in other isolated/custom partitions
func IAMServiceEC2(region string) string {
	partitions := endpoints.DefaultPartitions()
	for _, p := range partitions {
		if _, ok := p.Regions()[region]; ok {
			ep := "ec2." + p.DNSSuffix()
			return ep
		}
	}
	return "ec2.amazonaws.com"
}

// buildAWSIAMRolePolicy produces the AWS IAM role policy for the given role.
func (b *IAMModelBuilder) buildAWSIAMRolePolicy(role iam.Subject) (fi.Resource, error) {
	var policy string
	serviceAccount, ok := role.ServiceAccount()
	if ok {
		serviceAccountIssuer, err := iam.ServiceAccountIssuer(&b.Cluster.Spec)
		if err != nil {
			return nil, err
		}
		oidcProvider := strings.TrimPrefix(serviceAccountIssuer, "https://")

		iamPolicy := &iam.Policy{
			Version: iam.PolicyDefaultVersion,
			Statement: []*iam.Statement{
				{
					Effect: "Allow",
					Principal: iam.Principal{
						Federated: "arn:aws:iam::" + b.AWSAccountID + ":oidc-provider/" + oidcProvider,
					},
					Action: stringorslice.String("sts:AssumeRoleWithWebIdentity"),
					Condition: map[string]interface{}{
						"StringEquals": map[string]interface{}{
							oidcProvider + ":sub": "system:serviceaccount:" + serviceAccount.Namespace + ":" + serviceAccount.Name,
						},
					},
				},
			},
		}
		s, err := iamPolicy.AsJSON()
		if err != nil {
			return nil, err
		}
		policy = s
	} else {
		// We don't generate using json.Marshal here, it would create whitespace changes in the policy for existing clusters.

		policy = strings.ReplaceAll(NodeRolePolicyTemplate, "{{ IAMServiceEC2 }}", IAMServiceEC2(b.Region))
	}

	return fi.NewStringResource(policy), nil
}
