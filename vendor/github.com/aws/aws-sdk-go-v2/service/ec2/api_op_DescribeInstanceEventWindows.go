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

// Describes the specified event windows or all event windows.
//
// If you specify event window IDs, the output includes information for only the
// specified event windows. If you specify filters, the output includes information
// for only those event windows that meet the filter criteria. If you do not
// specify event windows IDs or filters, the output includes information for all
// event windows, which can affect performance. We recommend that you use
// pagination to ensure that the operation returns quickly and successfully.
//
// For more information, see [Define event windows for scheduled events] in the Amazon EC2 User Guide.
//
// [Define event windows for scheduled events]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html
func (c *Client) DescribeInstanceEventWindows(ctx context.Context, params *DescribeInstanceEventWindowsInput, optFns ...func(*Options)) (*DescribeInstanceEventWindowsOutput, error) {
	if params == nil {
		params = &DescribeInstanceEventWindowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceEventWindows", params, optFns, c.addOperationDescribeInstanceEventWindowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceEventWindowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Describe instance event windows by InstanceEventWindow.
type DescribeInstanceEventWindowsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//
	//   - dedicated-host-id - The event windows associated with the specified
	//   Dedicated Host ID.
	//
	//   - event-window-name - The event windows associated with the specified names.
	//
	//   - instance-id - The event windows associated with the specified instance ID.
	//
	//   - instance-tag - The event windows associated with the specified tag and value.
	//
	//   - instance-tag-key - The event windows associated with the specified tag key,
	//   regardless of the value.
	//
	//   - instance-tag-value - The event windows associated with the specified tag
	//   value, regardless of the key.
	//
	//   - tag: - The key/value combination of a tag assigned to the event window. Use
	//   the tag key in the filter name and the tag value as the filter value. For
	//   example, to find all resources that have a tag with the key Owner and the
	//   value CMX , specify tag:Owner for the filter name and CMX for the filter
	//   value.
	//
	//   - tag-key - The key of a tag assigned to the event window. Use this filter to
	//   find all event windows that have a tag with a specific key, regardless of the
	//   tag value.
	//
	//   - tag-value - The value of a tag assigned to the event window. Use this filter
	//   to find all event windows that have a tag with a specific value, regardless of
	//   the tag key.
	Filters []types.Filter

	// The IDs of the event windows.
	InstanceEventWindowIds []string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 20 and 500. You cannot specify this parameter and the event
	// window IDs parameter in the same call.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstanceEventWindowsOutput struct {

	// Information about the event windows.
	InstanceEventWindows []types.InstanceEventWindow

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceEventWindowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceEventWindows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceEventWindows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInstanceEventWindows"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceEventWindows(options.Region), middleware.Before); err != nil {
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

// DescribeInstanceEventWindowsAPIClient is a client that implements the
// DescribeInstanceEventWindows operation.
type DescribeInstanceEventWindowsAPIClient interface {
	DescribeInstanceEventWindows(context.Context, *DescribeInstanceEventWindowsInput, ...func(*Options)) (*DescribeInstanceEventWindowsOutput, error)
}

var _ DescribeInstanceEventWindowsAPIClient = (*Client)(nil)

// DescribeInstanceEventWindowsPaginatorOptions is the paginator options for
// DescribeInstanceEventWindows
type DescribeInstanceEventWindowsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 20 and 500. You cannot specify this parameter and the event
	// window IDs parameter in the same call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstanceEventWindowsPaginator is a paginator for
// DescribeInstanceEventWindows
type DescribeInstanceEventWindowsPaginator struct {
	options   DescribeInstanceEventWindowsPaginatorOptions
	client    DescribeInstanceEventWindowsAPIClient
	params    *DescribeInstanceEventWindowsInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstanceEventWindowsPaginator returns a new
// DescribeInstanceEventWindowsPaginator
func NewDescribeInstanceEventWindowsPaginator(client DescribeInstanceEventWindowsAPIClient, params *DescribeInstanceEventWindowsInput, optFns ...func(*DescribeInstanceEventWindowsPaginatorOptions)) *DescribeInstanceEventWindowsPaginator {
	if params == nil {
		params = &DescribeInstanceEventWindowsInput{}
	}

	options := DescribeInstanceEventWindowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstanceEventWindowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstanceEventWindowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstanceEventWindows page.
func (p *DescribeInstanceEventWindowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstanceEventWindowsOutput, error) {
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

	result, err := p.client.DescribeInstanceEventWindows(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInstanceEventWindows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInstanceEventWindows",
	}
}
