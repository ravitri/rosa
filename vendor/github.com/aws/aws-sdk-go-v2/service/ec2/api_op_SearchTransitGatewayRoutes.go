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

// Searches for routes in the specified transit gateway route table.
func (c *Client) SearchTransitGatewayRoutes(ctx context.Context, params *SearchTransitGatewayRoutesInput, optFns ...func(*Options)) (*SearchTransitGatewayRoutesOutput, error) {
	if params == nil {
		params = &SearchTransitGatewayRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchTransitGatewayRoutes", params, optFns, c.addOperationSearchTransitGatewayRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchTransitGatewayRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTransitGatewayRoutesInput struct {

	// One or more filters. The possible values are:
	//   - attachment.transit-gateway-attachment-id - The id of the transit gateway
	//   attachment.
	//   - attachment.resource-id - The resource id of the transit gateway attachment.
	//   - attachment.resource-type - The attachment resource type. Valid values are
	//   vpc | vpn | direct-connect-gateway | peering | connect .
	//   - prefix-list-id - The ID of the prefix list.
	//   - route-search.exact-match - The exact match of the specified filter.
	//   - route-search.longest-prefix-match - The longest prefix that matches the
	//   route.
	//   - route-search.subnet-of-match - The routes with a subnet that match the
	//   specified CIDR filter.
	//   - route-search.supernet-of-match - The routes with a CIDR that encompass the
	//   CIDR filter. For example, if you have 10.0.1.0/29 and 10.0.1.0/31 routes in your
	//   route table and you specify supernet-of-match as 10.0.1.0/30, then the result
	//   returns 10.0.1.0/29.
	//   - state - The state of the route ( active | blackhole ).
	//   - type - The type of route ( propagated | static ).
	//
	// This member is required.
	Filters []types.Filter

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum number of routes to return. If a value is not provided, the default
	// is 1000.
	MaxResults *int32

	noSmithyDocumentSerde
}

type SearchTransitGatewayRoutesOutput struct {

	// Indicates whether there are additional routes available.
	AdditionalRoutesAvailable *bool

	// Information about the routes.
	Routes []types.TransitGatewayRoute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchTransitGatewayRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpSearchTransitGatewayRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpSearchTransitGatewayRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchTransitGatewayRoutes"); err != nil {
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
	if err = addOpSearchTransitGatewayRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTransitGatewayRoutes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchTransitGatewayRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchTransitGatewayRoutes",
	}
}
