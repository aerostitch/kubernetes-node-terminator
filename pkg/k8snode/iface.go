package k8snode

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

type Node interface {
	status(metav1.ListOptions) (*v1.NodeList, error)
	terminate(v1.Node) error
}
