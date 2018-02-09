package main

import "testing"

import (
	"k8s.io/client-go/pkg/api/v1"
)

type FakeKubernetesClientConfig struct{}

var fakeNode = v1.Node{
	ObjectMeta: v1.ObjectMeta{},
	Spec:       v1.NodeSpec{},
}

func fakeNodeList(listOptions v1.ListOptions) *v1.NodeList {
	labels := make(map[string]string)
	labels["status"] = "unhealthy"

	if listOptions.LabelSelector != keysString(labels) {
		return &v1.NodeList{
			ListMeta: meta_v1.ListMeta{},
			Items:    []v1.Node{},
		}
	}

	fakeNode.Labels = labels

	nodeList := &v1.NodeList{
		ListMeta: meta_v1.ListMeta{},
		Items: []v1.Node{
			fakeNode,
		},
	}
	return nodeList
}

func newFakeClient() kubernetesClient {
	return &FakeKubernetesClientConfig{}
}

func (c FakeKubernetesClientConfig) getNodes(listOptions v1.ListOptions) (*v1.NodeList, error) {
	return fakeNodeList(listOptions), nil
}

func TestKubernetesNodes_GetNodesByLabel(t *testing.T) {
	client := newFakeClient()
	nodesController := kubernetesNodes{}
	labels := make(map[string]string)
	labels["status"] = "unhealthy"

	nodeList, err := nodesController.getNodesByLabel(client, labels)
	if err != nil {
		t.Errorf("failed to populate node by label: %s", err)
	}

	if len(nodeList.Items) <= 0 {
		t.Error("failed to get node by label")
	}

	for _, node := range nodeList.Items {
		if node.Labels["status"] != "unhealthy" {
			t.Errorf("expected unhealthy but got %s", node.Labels["status"])
		}
	}
}
