package main

import (
	"flag"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func newClient() (*kubernetes.Clientset, error) {
	var clientset *kubernetes.Clientset

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return clientset, err
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return clientset, err
	}
	return clientset, err
}

func main() {
	flag.Parse()

	flag.Lookup("logtostderr").Value.Set("true")

	glog.Info("Starting node-terminator")
	client, err := newClient()
	if err != nil {
		glog.Fatal(err.Error())
	}

	config := newConfig(client)

	labels := make(map[string]string)
	labels["status"] = "unhealthy"

	for {
		glog.Info("Checking for unhealthy instances")

		nodeList, err := config.status(labels)
		if err != nil {
			glog.Fatalf("failed to populate node by label: %s", err)
		}

		// iterate over nodeList and call config.status(<node>) if needed
	}
}
