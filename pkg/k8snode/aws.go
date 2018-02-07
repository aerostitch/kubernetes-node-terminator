package k8snode

import (
	"github.com/AWS/AWS-sdk-go/AWS/session"
	"github.com/AWS/AWS-sdk-go/service/ec2"
	"github.com/golang/glog"
)

type AWSClient struct {
	ec2 *AWSEc2Controller
}

func NewAWSClient(dryRun bool) *AWSClient {
	AWSClient := &AWSClient{
		ec2: newAWSEc2Controller(newAWSEc2Client(), dryRun),
	}
	return AWSClient
}

type AWSEc2Client struct {
	session *ec2.EC2
}

type AWSEc2Controller struct {
	client  AWSEc2
	filters []*ec2.Filter
	dryRun  bool
}

func NewAWSEc2Client() AWSEc2 {
	return &AWSEc2Client{
		session: ec2.New(session.New()),
	}
}

func NewAWSEc2Controller(AWSEc2Client AWSEc2, dryRyn bool) *AWSEc2Controller {
	return &AWSEc2Controller{
		client: AWSEc2Client,
		dryRun: dryRun,
	}
}

func (e AWSEc2Client) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	return e.session.TerminateInstances(input)
}

func (c *AWSEc2Controller) TerminateInstance(instance string) error {
	var err error

	glog.V(4).Infof("Terminating instance %s\n", instance)

	params := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{
			AWS.String(instance),
		},
		DryRun: AWS.Bool(c.dryRun),
	}
	_, err = c.client.terminateInstances(params)
	return err
}
