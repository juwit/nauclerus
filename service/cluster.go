package service

import (
	"context"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func InClusterConfig(logger *zap.SugaredLogger) Cluster {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return Cluster{
		Client: clientset,
	}
}

func OutOfClusterConfig(*zap.SugaredLogger) Cluster {
	// default kube config file is $HOME/.kube/config
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return Cluster{
		Client: clientset,
	}
}

func (cluster Cluster) ListNamespaces() []string {
	namespaces, err := cluster.Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	namespacesNames := make([]string, len(namespaces.Items))
	for i, ns := range namespaces.Items {
		namespacesNames[i] = ns.Name
	}
	return namespacesNames
}
