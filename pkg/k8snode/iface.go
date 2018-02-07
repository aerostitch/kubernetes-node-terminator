package k8snode

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

type Node interface {
	Status(metav1.ListOptions) (*v1.NodeList, error)
	Terminate(v1.Node) error
	Event(v1.Node) error
}

type Provider interface {
	TerminateInstance(string) error
}

/*type AWSec2 interface {
	terminateInstances(*ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)
}*/
