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

// Creates a Traffic Mirror filter rule. A Traffic Mirror rule defines the Traffic
// Mirror source traffic to mirror. You need the Traffic Mirror filter ID when you
// create the rule.
func (c *Client) CreateTrafficMirrorFilterRule(ctx context.Context, params *CreateTrafficMirrorFilterRuleInput, optFns ...func(*Options)) (*CreateTrafficMirrorFilterRuleOutput, error) {
	if params == nil {
		params = &CreateTrafficMirrorFilterRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrafficMirrorFilterRule", params, optFns, c.addOperationCreateTrafficMirrorFilterRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrafficMirrorFilterRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrafficMirrorFilterRuleInput struct {

	// The destination CIDR block to assign to the Traffic Mirror rule.
	//
	// This member is required.
	DestinationCidrBlock *string

	// The action to take on the filtered traffic.
	//
	// This member is required.
	RuleAction types.TrafficMirrorRuleAction

	// The number of the Traffic Mirror rule. This number must be unique for each
	// Traffic Mirror rule in a given direction. The rules are processed in ascending
	// order by rule number.
	//
	// This member is required.
	RuleNumber *int32

	// The source CIDR block to assign to the Traffic Mirror rule.
	//
	// This member is required.
	SourceCidrBlock *string

	// The type of traffic.
	//
	// This member is required.
	TrafficDirection types.TrafficDirection

	// The ID of the filter that this rule is associated with.
	//
	// This member is required.
	TrafficMirrorFilterId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// The description of the Traffic Mirror rule.
	Description *string

	// The destination port range.
	DestinationPortRange *types.TrafficMirrorPortRangeRequest

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The protocol, for example UDP, to assign to the Traffic Mirror rule. For
	// information about the protocol value, see Protocol Numbers (https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
	// on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *int32

	// The source port range.
	SourcePortRange *types.TrafficMirrorPortRangeRequest

	noSmithyDocumentSerde
}

type CreateTrafficMirrorFilterRuleOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// The Traffic Mirror rule.
	TrafficMirrorFilterRule *types.TrafficMirrorFilterRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrafficMirrorFilterRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTrafficMirrorFilterRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTrafficMirrorFilterRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTrafficMirrorFilterRule"); err != nil {
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
	if err = addIdempotencyToken_opCreateTrafficMirrorFilterRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTrafficMirrorFilterRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrafficMirrorFilterRule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateTrafficMirrorFilterRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTrafficMirrorFilterRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTrafficMirrorFilterRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTrafficMirrorFilterRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTrafficMirrorFilterRuleInput ")
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
func addIdempotencyToken_opCreateTrafficMirrorFilterRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateTrafficMirrorFilterRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTrafficMirrorFilterRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTrafficMirrorFilterRule",
	}
}
