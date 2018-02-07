package k8snode

import (
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

type Config struct {
	Node
	kclient  *kubernetes.Clientset
	provider *Provider
}

func NewConfig(kclient *kubernetes.Clientset, cloudType string, dryRun bool) Config {
	cfg := &Config{kclient: kclient}
	switch cloudType {
	case "aws":
		cfg.provider = &NewAWSClient(dryRun)
	default:
		glog.Fatalf("Cloud provider %s not supported\n", cloudType)
	}
}

func (c Config) Status(listOptions metav1.ListOptions) (*v1.NodeList, error) {
	nodeList, err := c.kclient.Core().Nodes().List(listOptions)
	return nodeList, err
}

func (c Config) Terminate(node v1.Node) error {
	i := node.Labels["instance-id"]
	o, err := c.provider.terminateInstance(i)
}

func (c Config) Event(node v1.node) error {
	// Create a kubernetes event for the given node object.   Something like "Terminating unhealthy instance"

	// placeholder to keep the corev1 import
	corev1.Event()
}
