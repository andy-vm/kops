// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of CIDR locations for the given collection (metadata
// only, does not include CIDR blocks).
func (c *Client) ListCidrLocations(ctx context.Context, params *ListCidrLocationsInput, optFns ...func(*Options)) (*ListCidrLocationsOutput, error) {
	if params == nil {
		params = &ListCidrLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCidrLocations", params, optFns, c.addOperationListCidrLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCidrLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCidrLocationsInput struct {

	// The CIDR collection ID.
	//
	// This member is required.
	CollectionId *string

	// The maximum number of CIDR collection locations to return in the response.
	MaxResults *int32

	// An opaque pagination token to indicate where the service is to begin
	// enumerating results.
	//
	// If no value is provided, the listing of results starts from the beginning.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCidrLocationsOutput struct {

	// A complex type that contains information about the list of CIDR locations.
	CidrLocations []types.LocationSummary

	// An opaque pagination token to indicate where the service is to begin
	// enumerating results.
	//
	// If no value is provided, the listing of results starts from the beginning.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCidrLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListCidrLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListCidrLocations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCidrLocations"); err != nil {
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
	if err = addOpListCidrLocationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCidrLocations(options.Region), middleware.Before); err != nil {
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

// ListCidrLocationsAPIClient is a client that implements the ListCidrLocations
// operation.
type ListCidrLocationsAPIClient interface {
	ListCidrLocations(context.Context, *ListCidrLocationsInput, ...func(*Options)) (*ListCidrLocationsOutput, error)
}

var _ ListCidrLocationsAPIClient = (*Client)(nil)

// ListCidrLocationsPaginatorOptions is the paginator options for ListCidrLocations
type ListCidrLocationsPaginatorOptions struct {
	// The maximum number of CIDR collection locations to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCidrLocationsPaginator is a paginator for ListCidrLocations
type ListCidrLocationsPaginator struct {
	options   ListCidrLocationsPaginatorOptions
	client    ListCidrLocationsAPIClient
	params    *ListCidrLocationsInput
	nextToken *string
	firstPage bool
}

// NewListCidrLocationsPaginator returns a new ListCidrLocationsPaginator
func NewListCidrLocationsPaginator(client ListCidrLocationsAPIClient, params *ListCidrLocationsInput, optFns ...func(*ListCidrLocationsPaginatorOptions)) *ListCidrLocationsPaginator {
	if params == nil {
		params = &ListCidrLocationsInput{}
	}

	options := ListCidrLocationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCidrLocationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCidrLocationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCidrLocations page.
func (p *ListCidrLocationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCidrLocationsOutput, error) {
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

	result, err := p.client.ListCidrLocations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCidrLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCidrLocations",
	}
}
