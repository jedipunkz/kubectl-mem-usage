package kube

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewClientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(config)
}

func GetNamespaces(clientset *kubernetes.Clientset) (*corev1.NamespaceList, error) {
	return clientset.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
}
