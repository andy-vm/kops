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

// For a specified resource ID, this API operation returns a list of compliance
// statuses for different resource types. Currently, you can only specify one
// resource ID per call. List results depend on the criteria specified in the
// filter.
func (c *Client) ListComplianceItems(ctx context.Context, params *ListComplianceItemsInput, optFns ...func(*Options)) (*ListComplianceItemsOutput, error) {
	if params == nil {
		params = &ListComplianceItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComplianceItems", params, optFns, c.addOperationListComplianceItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComplianceItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComplianceItemsInput struct {

	// One or more compliance filters. Use a filter to return a more specific list of
	// results.
	Filters []types.ComplianceStringFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// The ID for the resources from which to get compliance information. Currently,
	// you can only specify one resource ID.
	ResourceIds []string

	// The type of resource from which to get compliance information. Currently, the
	// only supported resource type is ManagedInstance .
	ResourceTypes []string

	noSmithyDocumentSerde
}

type ListComplianceItemsOutput struct {

	// A list of compliance information for the specified resource ID.
	ComplianceItems []types.ComplianceItem

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComplianceItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListComplianceItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListComplianceItems{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListComplianceItems"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComplianceItems(options.Region), middleware.Before); err != nil {
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

// ListComplianceItemsAPIClient is a client that implements the
// ListComplianceItems operation.
type ListComplianceItemsAPIClient interface {
	ListComplianceItems(context.Context, *ListComplianceItemsInput, ...func(*Options)) (*ListComplianceItemsOutput, error)
}

var _ ListComplianceItemsAPIClient = (*Client)(nil)

// ListComplianceItemsPaginatorOptions is the paginator options for
// ListComplianceItems
type ListComplianceItemsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComplianceItemsPaginator is a paginator for ListComplianceItems
type ListComplianceItemsPaginator struct {
	options   ListComplianceItemsPaginatorOptions
	client    ListComplianceItemsAPIClient
	params    *ListComplianceItemsInput
	nextToken *string
	firstPage bool
}

// NewListComplianceItemsPaginator returns a new ListComplianceItemsPaginator
func NewListComplianceItemsPaginator(client ListComplianceItemsAPIClient, params *ListComplianceItemsInput, optFns ...func(*ListComplianceItemsPaginatorOptions)) *ListComplianceItemsPaginator {
	if params == nil {
		params = &ListComplianceItemsInput{}
	}

	options := ListComplianceItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComplianceItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComplianceItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComplianceItems page.
func (p *ListComplianceItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComplianceItemsOutput, error) {
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

	result, err := p.client.ListComplianceItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComplianceItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListComplianceItems",
	}
}
