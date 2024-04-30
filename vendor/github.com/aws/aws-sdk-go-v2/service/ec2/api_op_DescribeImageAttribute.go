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

// Describes the specified attribute of the specified AMI. You can specify only
// one attribute at a time. The order of the elements in the response, including
// those within nested structures, might vary. Applications should not assume the
// elements appear in a particular order.
func (c *Client) DescribeImageAttribute(ctx context.Context, params *DescribeImageAttributeInput, optFns ...func(*Options)) (*DescribeImageAttributeOutput, error) {
	if params == nil {
		params = &DescribeImageAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImageAttribute", params, optFns, c.addOperationDescribeImageAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImageAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeImageAttribute.
type DescribeImageAttributeInput struct {

	// The AMI attribute. Note: The blockDeviceMapping attribute is deprecated. Using
	// this attribute returns the Client.AuthFailure error. To get information about
	// the block device mappings for an AMI, use the DescribeImages action.
	//
	// This member is required.
	Attribute types.ImageAttributeName

	// The ID of the AMI.
	//
	// This member is required.
	ImageId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

// Describes an image attribute.
type DescribeImageAttributeOutput struct {

	// The block device mapping entries.
	BlockDeviceMappings []types.BlockDeviceMapping

	// The boot mode.
	BootMode *types.AttributeValue

	// A description for the AMI.
	Description *types.AttributeValue

	// The ID of the AMI.
	ImageId *string

	// If v2.0 , it indicates that IMDSv2 is specified in the AMI. Instances launched
	// from this AMI will have HttpTokens automatically set to required so that, by
	// default, the instance requires that IMDSv2 is used when requesting instance
	// metadata. In addition, HttpPutResponseHopLimit is set to 2 . For more
	// information, see Configure the AMI (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration)
	// in the Amazon EC2 User Guide.
	ImdsSupport *types.AttributeValue

	// The kernel ID.
	KernelId *types.AttributeValue

	// The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601)
	// , when the AMI was last used to launch an EC2 instance. When the AMI is used to
	// launch an instance, there is a 24-hour delay before that usage is reported.
	// lastLaunchedTime data is available starting April 2017.
	LastLaunchedTime *types.AttributeValue

	// The launch permissions.
	LaunchPermissions []types.LaunchPermission

	// The product codes.
	ProductCodes []types.ProductCode

	// The RAM disk ID.
	RamdiskId *types.AttributeValue

	// Indicates whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport *types.AttributeValue

	// If the image is configured for NitroTPM support, the value is v2.0 .
	TpmSupport *types.AttributeValue

	// Base64 representation of the non-volatile UEFI variable store. To retrieve the
	// UEFI data, use the GetInstanceUefiData (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceUefiData)
	// command. You can inspect and modify the UEFI data by using the python-uefivars
	// tool (https://github.com/awslabs/python-uefivars) on GitHub. For more
	// information, see UEFI Secure Boot (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-secure-boot.html)
	// in the Amazon EC2 User Guide.
	UefiData *types.AttributeValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImageAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeImageAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeImageAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeImageAttribute"); err != nil {
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
	if err = addOpDescribeImageAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeImageAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeImageAttribute",
	}
}
