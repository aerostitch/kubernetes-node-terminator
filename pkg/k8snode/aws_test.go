package k8snode

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type MockAWSEc2Controller struct {
	client  ec2iface.EC2API
	filters []*ec2.Filter
	dryRun  bool
}

func NewMockAWSEc2Controller(dryRun bool, region string) Provider {

	return &MockAWSEc2Controller{
		client: ec2.New(session.New(&aws.Config{Region: aws.String(region)})),
		dryRun: dryRun,
	}
}

func TestTerminateInstances(t *testing.T) {
	mockSvc := NewMockAWSEc2Controller(false, "us-east-1")
	instanceID := "i-0ed48177c77a0acfb"

}
