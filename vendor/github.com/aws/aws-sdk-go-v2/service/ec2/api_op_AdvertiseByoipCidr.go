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

// Advertises an IPv4 or IPv6 address range that is provisioned for use with your
// Amazon Web Services resources through bring your own IP addresses (BYOIP).
//
// You can perform this operation at most once every 10 seconds, even if you
// specify different address ranges each time.
//
// We recommend that you stop advertising the BYOIP CIDR from other locations when
// you advertise it from Amazon Web Services. To minimize down time, you can
// configure your Amazon Web Services resources to use an address from a BYOIP CIDR
// before it is advertised, and then simultaneously stop advertising it from the
// current location and start advertising it through Amazon Web Services.
//
// It can take a few minutes before traffic to the specified addresses starts
// routing to Amazon Web Services because of BGP propagation delays.
//
// To stop advertising the BYOIP CIDR, use WithdrawByoipCidr.
func (c *Client) AdvertiseByoipCidr(ctx context.Context, params *AdvertiseByoipCidrInput, optFns ...func(*Options)) (*AdvertiseByoipCidrOutput, error) {
	if params == nil {
		params = &AdvertiseByoipCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdvertiseByoipCidr", params, optFns, c.addOperationAdvertiseByoipCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdvertiseByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdvertiseByoipCidrInput struct {

	// The address range, in CIDR notation. This must be the exact range that you
	// provisioned. You can't advertise only a portion of the provisioned range.
	//
	// This member is required.
	Cidr *string

	// The public 2-byte or 4-byte ASN that you want to advertise.
	Asn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// If you have [Local Zones] enabled, you can choose a network border group for Local Zones
	// when you provision and advertise a BYOIPv4 CIDR. Choose the network border group
	// carefully as the EIP and the Amazon Web Services resource it is associated with
	// must reside in the same network border group.
	//
	// You can provision BYOIP address ranges to and advertise them in the following
	// Local Zone network border groups:
	//
	//   - us-east-1-dfw-2
	//
	//   - us-west-2-lax-1
	//
	//   - us-west-2-phx-2
	//
	// You cannot provision or advertise BYOIPv6 address ranges in Local Zones at this
	// time.
	//
	// [Local Zones]: https://docs.aws.amazon.com/local-zones/latest/ug/how-local-zones-work.html
	NetworkBorderGroup *string

	noSmithyDocumentSerde
}

type AdvertiseByoipCidrOutput struct {

	// Information about the address range.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdvertiseByoipCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpAdvertiseByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAdvertiseByoipCidr{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdvertiseByoipCidr"); err != nil {
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
	if err = addOpAdvertiseByoipCidrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdvertiseByoipCidr(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdvertiseByoipCidr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdvertiseByoipCidr",
	}
}
