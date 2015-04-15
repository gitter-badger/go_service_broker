package aws_client

import (
	"strconv"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/ec2"
)

const (
	AMI_ID            = "ami-80280ee8"
	SECURITY_GROUP_ID = "sg-b23aead6"
	SUBNET_ID         = "subnet-0c75a427"
	KEY_NAME          = "mykey1"
	INSTANCE_TYPE     = "t2.micro"
)

type Client interface {
	CreateInstance(imageId string) (string, error)
	GetInstanceState(instancId string) (string, error)
	CreateKeyPair(keyName string) (string, error)
}

type AWSClient struct {
	EC2Client *ec2.EC2
}

func NewClient(region string) *AWSClient {
	return &AWSClient{
		EC2Client: ec2.New(&aws.Config{Region: region}),
	}
}

func (c *AWSClient) CreateInstance() (string, error) {
	return c.createInstance(AMI_ID)
}

func (c *AWSClient) GetInstanceState(instanceId string) (string, error) {
	instanceInput := &ec2.DescribeInstancesInput{
		InstanceIDs: []*string{
			aws.String(instanceId), // Required
		},
	}

	instanceOutput, err := c.EC2Client.DescribeInstances(instanceInput)
	if err != nil {
		return "", err
	}

	state, _ := strconv.Unquote(awsutil.StringValue(instanceOutput.Reservations[0].Instances[0].State.Name))
	return state, nil
}

func (c *AWSClient) CreateKeyPair(keyName string) (string, error) {
	keypairInput := &ec2.CreateKeyPairInput{
		KeyName: aws.String(keyName),
	}

	keypairOutput, err := c.EC2Client.CreateKeyPair(keypairInput)
	if err != nil {
		return "", err
	}
	return awsutil.StringValue(keypairOutput), nil
}

func (c *AWSClient) createInstance(imageId string) (string, error) {
	instanceInput := &ec2.RunInstancesInput{
		ImageID:  aws.String(imageId), // Required
		MaxCount: aws.Long(1),         // Required
		MinCount: aws.Long(1),         // Required
		// AdditionalInfo: aws.String("String"),
		// BlockDeviceMappings: []*ec2.BlockDeviceMapping{
		// 	&ec2.BlockDeviceMapping{ // Required
		// 		DeviceName: aws.String("String"),
		// 		EBS: &ec2.EBSBlockDevice{
		// 			DeleteOnTermination: aws.Boolean(true),
		// 			Encrypted:           aws.Boolean(true),
		// 			IOPS:                aws.Long(1),
		// 			SnapshotID:          aws.String("String"),
		// 			VolumeSize:          aws.Long(1),
		// 			VolumeType:          aws.String("VolumeType"),
		// 		},
		// 		NoDevice:    aws.String("String"),
		// 		VirtualName: aws.String("String"),
		// 	},
		// 	// More values...
		// },
		// ClientToken: aws.String("String"),
		// DisableAPITermination: aws.Boolean(true),
		// DryRun:                aws.Boolean(true),
		// EBSOptimized:          aws.Boolean(true),
		// IAMInstanceProfile: &ec2.IAMInstanceProfileSpecification{
		// 	ARN:  aws.String("String"),
		// 	Name: aws.String("String"),
		// },
		// InstanceInitiatedShutdownBehavior: aws.String("ShutdownBehavior"),
		InstanceType: aws.String(INSTANCE_TYPE),
		// KernelID:                          aws.String("String"),
		KeyName: aws.String(KEY_NAME),
		// Monitoring: &ec2.RunInstancesMonitoringEnabled{
		// 	Enabled: aws.Boolean(true), // Required
		// },
		// NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
		// 	&ec2.InstanceNetworkInterfaceSpecification{ // Required
		// 		AssociatePublicIPAddress: aws.Boolean(true),
		// 		DeleteOnTermination:      aws.Boolean(true),
		// 		Description:              aws.String("String"),
		// 		DeviceIndex:              aws.Long(1),
		// 		Groups: []*string{
		// 			aws.String("String"), // Required
		// 			// More values...
		// 		},
		// 		NetworkInterfaceID: aws.String("String"),
		// 		PrivateIPAddress:   aws.String("String"),
		// 		PrivateIPAddresses: []*ec2.PrivateIPAddressSpecification{
		// 			&ec2.PrivateIPAddressSpecification{ // Required
		// 				PrivateIPAddress: aws.String("String"), // Required
		// 				Primary:          aws.Boolean(true),
		// 			},
		// 			// More values...
		// 		},
		// 		SecondaryPrivateIPAddressCount: aws.Long(1),
		// 		SubnetID:                       aws.String("String"),
		// 	},
		// 	// More values...
		// },
		// Placement: &ec2.Placement{
		// 	AvailabilityZone: aws.String("String"),
		// 	GroupName:        aws.String("String"),
		// 	Tenancy:          aws.String("Tenancy"),
		// },
		// PrivateIPAddress: aws.String("String"),
		// RAMDiskID:        aws.String("String"),
		SecurityGroupIDs: []*string{
			aws.String(SECURITY_GROUP_ID), // Required
			// More values...
		},
		SubnetID: aws.String(SUBNET_ID),
	}

	instancOutput, err := c.EC2Client.RunInstances(instanceInput)
	if err != nil {
		return "", err
	}

	instanceId, _ := strconv.Unquote(awsutil.StringValue(instancOutput.Instances[0].InstanceID))

	return instanceId, nil
}