// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// An Amazon Web Services Verified Access endpoint is where you define your
// application along with an optional endpoint-level access policy.
func (c *Client) CreateVerifiedAccessEndpoint(ctx context.Context, params *CreateVerifiedAccessEndpointInput, optFns ...func(*Options)) (*CreateVerifiedAccessEndpointOutput, error) {
	if params == nil {
		params = &CreateVerifiedAccessEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVerifiedAccessEndpoint", params, optFns, c.addOperationCreateVerifiedAccessEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVerifiedAccessEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVerifiedAccessEndpointInput struct {

	// The DNS name for users to reach your application.
	//
	// This member is required.
	ApplicationDomain *string

	// The type of attachment.
	//
	// This member is required.
	AttachmentType types.VerifiedAccessEndpointAttachmentType

	// The ARN of the public TLS/SSL certificate in Amazon Web Services Certificate
	// Manager to associate with the endpoint. The CN in the certificate must match the
	// DNS name your end users will use to reach your application.
	//
	// This member is required.
	DomainCertificateArn *string

	// A custom identifier that is prepended to the DNS name that is generated for the
	// endpoint.
	//
	// This member is required.
	EndpointDomainPrefix *string

	// The type of Verified Access endpoint to create.
	//
	// This member is required.
	EndpointType types.VerifiedAccessEndpointType

	// The ID of the Verified Access group to associate the endpoint with.
	//
	// This member is required.
	VerifiedAccessGroupId *string

	// A unique, case-sensitive token that you provide to ensure idempotency of your
	// modification request. For more information, see [Ensuring Idempotency].
	//
	// [Ensuring Idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// A description for the Verified Access endpoint.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The load balancer details. This parameter is required if the endpoint type is
	// load-balancer .
	LoadBalancerOptions *types.CreateVerifiedAccessEndpointLoadBalancerOptions

	// The network interface details. This parameter is required if the endpoint type
	// is network-interface .
	NetworkInterfaceOptions *types.CreateVerifiedAccessEndpointEniOptions

	// The Verified Access policy document.
	PolicyDocument *string

	// The IDs of the security groups to associate with the Verified Access endpoint.
	// Required if AttachmentType is set to vpc .
	SecurityGroupIds []string

	// The options for server side encryption.
	SseSpecification *types.VerifiedAccessSseSpecificationRequest

	// The tags to assign to the Verified Access endpoint.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateVerifiedAccessEndpointOutput struct {

	// Details about the Verified Access endpoint.
	VerifiedAccessEndpoint *types.VerifiedAccessEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVerifiedAccessEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVerifiedAccessEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVerifiedAccessEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVerifiedAccessEndpoint"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateVerifiedAccessEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateVerifiedAccessEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVerifiedAccessEndpoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateVerifiedAccessEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateVerifiedAccessEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateVerifiedAccessEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateVerifiedAccessEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateVerifiedAccessEndpointInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateVerifiedAccessEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateVerifiedAccessEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateVerifiedAccessEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVerifiedAccessEndpoint",
	}
}
