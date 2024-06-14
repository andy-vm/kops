// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all patch groups that have been registered with patch baselines.
func (c *Client) DescribePatchGroups(ctx context.Context, params *DescribePatchGroupsInput, optFns ...func(*Options)) (*DescribePatchGroupsOutput, error) {
	if params == nil {
		params = &DescribePatchGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePatchGroups", params, optFns, c.addOperationDescribePatchGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePatchGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePatchGroupsInput struct {

	// Each element in the array is a structure containing a key-value pair.
	//
	// Supported keys for DescribePatchGroups include the following:
	//
	//   - NAME_PREFIX
	//
	// Sample values: AWS- | My- .
	//
	//   - OPERATING_SYSTEM
	//
	// Sample values: AMAZON_LINUX | SUSE | WINDOWS
	Filters []types.PatchOrchestratorFilter

	// The maximum number of patch groups to return (per page).
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribePatchGroupsOutput struct {

	// Each entry in the array contains:
	//
	//   - PatchGroup : string (between 1 and 256 characters. Regex:
	//   ^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$)
	//
	//   - PatchBaselineIdentity : A PatchBaselineIdentity element.
	Mappings []types.PatchGroupPatchBaselineMapping

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePatchGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePatchGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePatchGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePatchGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePatchGroups(options.Region), middleware.Before); err != nil {
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

// DescribePatchGroupsAPIClient is a client that implements the
// DescribePatchGroups operation.
type DescribePatchGroupsAPIClient interface {
	DescribePatchGroups(context.Context, *DescribePatchGroupsInput, ...func(*Options)) (*DescribePatchGroupsOutput, error)
}

var _ DescribePatchGroupsAPIClient = (*Client)(nil)

// DescribePatchGroupsPaginatorOptions is the paginator options for
// DescribePatchGroups
type DescribePatchGroupsPaginatorOptions struct {
	// The maximum number of patch groups to return (per page).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePatchGroupsPaginator is a paginator for DescribePatchGroups
type DescribePatchGroupsPaginator struct {
	options   DescribePatchGroupsPaginatorOptions
	client    DescribePatchGroupsAPIClient
	params    *DescribePatchGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribePatchGroupsPaginator returns a new DescribePatchGroupsPaginator
func NewDescribePatchGroupsPaginator(client DescribePatchGroupsAPIClient, params *DescribePatchGroupsInput, optFns ...func(*DescribePatchGroupsPaginatorOptions)) *DescribePatchGroupsPaginator {
	if params == nil {
		params = &DescribePatchGroupsInput{}
	}

	options := DescribePatchGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePatchGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePatchGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePatchGroups page.
func (p *DescribePatchGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePatchGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribePatchGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribePatchGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePatchGroups",
	}
}
