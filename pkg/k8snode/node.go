package k8snode

import (
	"fmt"
	"strings"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Config struct {
	kclient  *kubernetes.Clientset
	provider Provider
}

func NewConfig(kclient *kubernetes.Clientset, cloudType string, cloudRegion string, dryRun bool) Node {
	cfg := &Config{kclient: kclient}
	switch cloudType {
	case "aws":
		cfg.provider = NewAWSEc2Controller(dryRun, cloudRegion)
	default:
		glog.Fatalf("Cloud provider %s not supported\n", cloudType)
	}
	return cfg
}

func (c Config) Status(labels map[string]string) (*corev1.NodeList, error) {
	nodeList, err := c.kclient.Core().Nodes().List(c.labelsToListOptions(labels))
	return nodeList, err
}

func (c Config) Terminate(instanceID string) error {
	err := c.provider.TerminateInstance(instanceID)
	return err
}

func (c Config) Event(node corev1.Node) error {
	// Create a kubernetes event for the given node object.   Something like "Terminating unhealthy instance"

	// placeholder to keep the corev1 import
	//corev1.Event()
	return nil
}

func (c Config) labelsToListOptions(labels map[string]string) metav1.ListOptions {
	keys := make([]string, 0, len(labels))
	for k, v := range labels {
		keys = append(keys, fmt.Sprintf("%s=%s", k, v))
	}

	return metav1.ListOptions{
		LabelSelector: strings.Join(keys, ","),
	}
}
