package k8snode

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

type Node interface {
	Status(metav1.ListOptions) (*v1.NodeList, error)
	Terminate(v1.Node) error
}

type AWSec2 interface {
	terminateInstances(*ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)
}
