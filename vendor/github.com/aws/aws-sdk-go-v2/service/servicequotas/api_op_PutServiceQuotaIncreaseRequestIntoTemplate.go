// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a quota increase request to your quota request template.
func (c *Client) PutServiceQuotaIncreaseRequestIntoTemplate(ctx context.Context, params *PutServiceQuotaIncreaseRequestIntoTemplateInput, optFns ...func(*Options)) (*PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
	if params == nil {
		params = &PutServiceQuotaIncreaseRequestIntoTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutServiceQuotaIncreaseRequestIntoTemplate", params, optFns, c.addOperationPutServiceQuotaIncreaseRequestIntoTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutServiceQuotaIncreaseRequestIntoTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutServiceQuotaIncreaseRequestIntoTemplateInput struct {

	// Specifies the Amazon Web Services Region to which the template applies.
	//
	// This member is required.
	AwsRegion *string

	// Specifies the new, increased value for the quota.
	//
	// This member is required.
	DesiredValue *float64

	// Specifies the quota identifier. To find the quota code for a specific quota,
	// use the ListServiceQuotas operation, and look for the QuotaCode response in the
	// output for the quota you want.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServices operation.
	//
	// This member is required.
	ServiceCode *string

	noSmithyDocumentSerde
}

type PutServiceQuotaIncreaseRequestIntoTemplateOutput struct {

	// Information about the quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutServiceQuotaIncreaseRequestIntoTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutServiceQuotaIncreaseRequestIntoTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutServiceQuotaIncreaseRequestIntoTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutServiceQuotaIncreaseRequestIntoTemplate"); err != nil {
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
	if err = addOpPutServiceQuotaIncreaseRequestIntoTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutServiceQuotaIncreaseRequestIntoTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutServiceQuotaIncreaseRequestIntoTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutServiceQuotaIncreaseRequestIntoTemplate",
	}
}
