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

// Removes an ingress authorization rule from a Client VPN endpoint.
func (c *Client) RevokeClientVpnIngress(ctx context.Context, params *RevokeClientVpnIngressInput, optFns ...func(*Options)) (*RevokeClientVpnIngressOutput, error) {
	if params == nil {
		params = &RevokeClientVpnIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeClientVpnIngress", params, optFns, c.addOperationRevokeClientVpnIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeClientVpnIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeClientVpnIngressInput struct {

	// The ID of the Client VPN endpoint with which the authorization rule is
	// associated.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The IPv4 address range, in CIDR notation, of the network for which access is
	// being removed.
	//
	// This member is required.
	TargetNetworkCidr *string

	// The ID of the Active Directory group for which to revoke access.
	AccessGroupId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Indicates whether access should be revoked for all clients.
	RevokeAllGroups *bool

	noSmithyDocumentSerde
}

type RevokeClientVpnIngressOutput struct {

	// The current state of the authorization rule.
	Status *types.ClientVpnAuthorizationRuleStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeClientVpnIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpRevokeClientVpnIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRevokeClientVpnIngress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RevokeClientVpnIngress"); err != nil {
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
	if err = addOpRevokeClientVpnIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeClientVpnIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeClientVpnIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RevokeClientVpnIngress",
	}
}
