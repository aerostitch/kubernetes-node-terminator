package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/VEVO/kubernetes-node-terminator/pkg/k8snode"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	cloudProvider = os.Getenv("CLOUD_PROVIDER")
	dryRun        = false
	dryRunStr     = os.Getenv("DRY_RUN")
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
	var err error
	flag.Parse()

	flag.Lookup("logtostderr").Value.Set("true")

	switch {
	case dryRunStr != "":
		dryRun, err = strconv.ParseBool(dryRunStr)
		if err != nil {
			glog.Fatalf("Error parsing DRY_RUN value: %s", err)
		}
	case cloudProvider == "":
		glog.Fatal("Set the CLOUD_PROVIDER variable")
	}

	glog.Info("Starting node-terminator")
	client, err := newClient()
	if err != nil {
		glog.Fatal(err.Error())
	}

	config := k8snode.NewConfig(client, cloudProvider, dryRun)

	labels := make(map[string]string)
	labels["status"] = "unhealthy"

	for {
		glog.Info("Checking for unhealthy instances")

		nodeList, err := config.Status(labels)
		if err != nil {
			glog.Fatalf("failed to populate node by label: %s", err)
		}

		for _, i := range nodeList.Items {
			instanceID := i.Labels["instance-id"]
			glog.Infof("InstanceID is %s\n", instanceID)
		}
	}
}
