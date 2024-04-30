package aws_client

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/openshift-online/ocm-common/pkg/log"

	elb "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

type AWSClient struct {
	Ec2Client            *ec2.Client
	Route53Client        *route53.Client
	StackFormationClient *cloudformation.Client
	ElbClient            *elb.Client
	StsClient            *sts.Client
	Region               string
	IamClient            *iam.Client
	ClientContext        context.Context
	AccountID            string
	KmsClient            *kms.Client
	CloudWatchLogsClient *cloudwatchlogs.Client
}

func CreateAWSClient(profileName string, region string) (*AWSClient, error) {
	var cfg aws.Config
	var err error

	if envCredential() {
		log.LogInfo("Got AWS_ACCESS_KEY_ID env settings, going to build the config with the env")
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					os.Getenv("AWS_ACCESS_KEY_ID"),
					os.Getenv("AWS_SECRET_ACCESS_KEY"),
					"")),
		)
	} else {
		if envAwsProfile() {
			file := os.Getenv("AWS_SHARED_CREDENTIALS_FILE")
			log.LogInfo("Got file path: %s from env variable AWS_SHARED_CREDENTIALS_FILE\n", file)
			cfg, err = config.LoadDefaultConfig(context.TODO(),
				config.WithRegion(region),
				config.WithSharedCredentialsFiles([]string{file}),
			)
		} else {
			cfg, err = config.LoadDefaultConfig(context.TODO(),
				config.WithRegion(region),
				config.WithSharedConfigProfile(profileName),
			)
		}

	}

	if err != nil {
		return nil, err
	}

	awsClient := &AWSClient{
		Ec2Client:            ec2.NewFromConfig(cfg),
		Route53Client:        route53.NewFromConfig(cfg),
		StackFormationClient: cloudformation.NewFromConfig(cfg),
		ElbClient:            elb.NewFromConfig(cfg),
		Region:               region,
		StsClient:            sts.NewFromConfig(cfg),
		IamClient:            iam.NewFromConfig(cfg),
		ClientContext:        context.TODO(),
		KmsClient:            kms.NewFromConfig(cfg),
	}
	awsClient.AccountID = awsClient.GetAWSAccountID()
	return awsClient, nil
}

func (client *AWSClient) GetAWSAccountID() string {
	input := &sts.GetCallerIdentityInput{}
	out, err := client.StsClient.GetCallerIdentity(client.ClientContext, input)
	if err != nil {
		return ""
	}
	return *out.Account
}

func (client *AWSClient) EC2() *ec2.Client {
	return client.Ec2Client
}

func (client *AWSClient) Route53() *route53.Client {
	return client.Route53Client
}
func (client *AWSClient) CloudFormation() *cloudformation.Client {
	return client.StackFormationClient
}
func (client *AWSClient) ELB() *elb.Client {
	return client.ElbClient
}
