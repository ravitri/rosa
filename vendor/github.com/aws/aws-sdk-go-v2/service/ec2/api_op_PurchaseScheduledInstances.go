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

// You can no longer purchase Scheduled Instances. Purchases the Scheduled
// Instances with the specified schedule. Scheduled Instances enable you to
// purchase Amazon EC2 compute capacity by the hour for a one-year term. Before you
// can purchase a Scheduled Instance, you must call
// DescribeScheduledInstanceAvailability to check for available schedules and
// obtain a purchase token. After you purchase a Scheduled Instance, you must call
// RunScheduledInstances during each scheduled time period. After you purchase a
// Scheduled Instance, you can't cancel, modify, or resell your purchase.
func (c *Client) PurchaseScheduledInstances(ctx context.Context, params *PurchaseScheduledInstancesInput, optFns ...func(*Options)) (*PurchaseScheduledInstancesOutput, error) {
	if params == nil {
		params = &PurchaseScheduledInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseScheduledInstances", params, optFns, c.addOperationPurchaseScheduledInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseScheduledInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for PurchaseScheduledInstances.
type PurchaseScheduledInstancesInput struct {

	// The purchase requests.
	//
	// This member is required.
	PurchaseRequests []types.PurchaseRequest

	// Unique, case-sensitive identifier that ensures the idempotency of the request.
	// For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

// Contains the output of PurchaseScheduledInstances.
type PurchaseScheduledInstancesOutput struct {

	// Information about the Scheduled Instances.
	ScheduledInstanceSet []types.ScheduledInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPurchaseScheduledInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpPurchaseScheduledInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpPurchaseScheduledInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PurchaseScheduledInstances"); err != nil {
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
	if err = addIdempotencyToken_opPurchaseScheduledInstancesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPurchaseScheduledInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseScheduledInstances(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPurchaseScheduledInstances struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPurchaseScheduledInstances) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPurchaseScheduledInstances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PurchaseScheduledInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PurchaseScheduledInstancesInput ")
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
func addIdempotencyToken_opPurchaseScheduledInstancesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPurchaseScheduledInstances{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPurchaseScheduledInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PurchaseScheduledInstances",
	}
}
