package k8snode

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

type Node interface {
	status(metav1.ListOptions) (*v1.NodeList, error)
	terminate(v1.Node) error
}

type Config struct {
	kclient *kubernetes.Clientset
}

func newConfig(kclient *kubernetes.Clientset) Config {
	return &Config{kclient: kclient}
}

func (c Config) status(listOptions metav1.ListOptions) (*v1.NodeList, error) {
	nodeList, err := c.kclient.Core().Nodes().List(listOptions)
	return nodeList, err
}

func (c Config) terminate(node v1.Node) error {
	/*  I wasn't sure the best place to put the ec2 code.   Should I create a seperate package for the ec2 code and then when ppl call the newConfig function they need to pass in the awsclient object as well?   So then the Config struct would look like this isntead

	type Config struct {
		kclient *kubernetes.Clientset
		awsclient *awsEc2Controller
	}
	*/
}

func (c Config) event(node v1.node) error {
	// Create a kubernetes event for the given node object.   Something like "Terminating unhealthy instance"

	// placeholder to keep the corev1 import
	corev1.Event()
}
