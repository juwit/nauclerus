package service

import (
	"k8s.io/client-go/kubernetes"
)

type GreetService func(...string) (string, error)

type Cluster struct {
	Client *kubernetes.Clientset
}
