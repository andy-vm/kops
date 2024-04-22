// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about all completed key material rotations for the
// specified KMS key. You must specify the KMS key in all requests. You can refine
// the key rotations list by limiting the number of rotations returned. For
// detailed information about automatic and on-demand key rotations, see Rotating
// KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
// in the Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a KMS key in a different Amazon Web Services account.
// Required permissions: kms:ListKeyRotations (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//   - EnableKeyRotation
//   - DisableKeyRotation
//   - GetKeyRotationStatus
//   - RotateKeyOnDemand
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see KMS eventual consistency (https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html)
// .
func (c *Client) ListKeyRotations(ctx context.Context, params *ListKeyRotationsInput, optFns ...func(*Options)) (*ListKeyRotationsOutput, error) {
	if params == nil {
		params = &ListKeyRotationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKeyRotations", params, optFns, c.addOperationListKeyRotationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKeyRotationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKeyRotationsInput struct {

	// Gets the key rotations for the specified KMS key. Specify the key ID or key ARN
	// of the KMS key. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey .
	//
	// This member is required.
	KeyId *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, KMS does not return more than the specified number of items,
	// but it might return fewer. This value is optional. If you include a value, it
	// must be between 1 and 1000, inclusive. If you do not include a value, it
	// defaults to 100.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated
	// response you just received.
	Marker *string

	noSmithyDocumentSerde
}

type ListKeyRotationsOutput struct {

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request.
	NextMarker *string

	// A list of completed key material rotations.
	Rotations []types.RotationsListEntry

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in this response to the Marker parameter in a
	// subsequent request.
	Truncated bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKeyRotationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListKeyRotations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListKeyRotations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKeyRotations"); err != nil {
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
	if err = addOpListKeyRotationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKeyRotations(options.Region), middleware.Before); err != nil {
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

// ListKeyRotationsAPIClient is a client that implements the ListKeyRotations
// operation.
type ListKeyRotationsAPIClient interface {
	ListKeyRotations(context.Context, *ListKeyRotationsInput, ...func(*Options)) (*ListKeyRotationsOutput, error)
}

var _ ListKeyRotationsAPIClient = (*Client)(nil)

// ListKeyRotationsPaginatorOptions is the paginator options for ListKeyRotations
type ListKeyRotationsPaginatorOptions struct {
	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, KMS does not return more than the specified number of items,
	// but it might return fewer. This value is optional. If you include a value, it
	// must be between 1 and 1000, inclusive. If you do not include a value, it
	// defaults to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKeyRotationsPaginator is a paginator for ListKeyRotations
type ListKeyRotationsPaginator struct {
	options   ListKeyRotationsPaginatorOptions
	client    ListKeyRotationsAPIClient
	params    *ListKeyRotationsInput
	nextToken *string
	firstPage bool
}

// NewListKeyRotationsPaginator returns a new ListKeyRotationsPaginator
func NewListKeyRotationsPaginator(client ListKeyRotationsAPIClient, params *ListKeyRotationsInput, optFns ...func(*ListKeyRotationsPaginatorOptions)) *ListKeyRotationsPaginator {
	if params == nil {
		params = &ListKeyRotationsInput{}
	}

	options := ListKeyRotationsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKeyRotationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKeyRotationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKeyRotations page.
func (p *ListKeyRotationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKeyRotationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListKeyRotations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListKeyRotations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKeyRotations",
	}
}
