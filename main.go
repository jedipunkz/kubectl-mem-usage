package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func getPodMetrics(clientset *metricsv.Clientset, namespace string) (*metricsv1beta1.PodMetricsList, error) {
	return clientset.MetricsV1beta1().PodMetricses(namespace).List(context.Background(), v1.ListOptions{})
}

func prepareBarData(podMetrics *metricsv1beta1.PodMetricsList) []pterm.Bar {
	bars := []pterm.Bar{}
	for _, podMetric := range podMetrics.Items {
		for _, container := range podMetric.Containers {
			memoryUsage := container.Usage["memory"]
			bars = append(bars, pterm.Bar{Label: container.Name, Value: int(memoryUsage.Value())})
		}
	}
	return bars
}

func getNamespaces(clientset *kubernetes.Clientset) (*corev1.NamespaceList, error) {
	return clientset.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
}

func main() {
	// Get current user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Construct path to kubeconfig file
	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	metricsClientset, err := metricsv.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// List all namespaces
	namespaces, err := getNamespaces(clientset)
	if err != nil {
		panic(err)
	}

	bars := []pterm.Bar{}
	for _, namespace := range namespaces.Items {
		podMetrics, err := getPodMetrics(metricsClientset, namespace.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		bars = append(bars, prepareBarData(podMetrics)...)
	}

	pterm.DefaultBarChart.WithBars(bars).WithHorizontal().WithWidth(50).Render()
}