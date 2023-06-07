package metrics

import (
	"context"

	"github.com/jedipunkz/kubectl-mem-usage/internal/kube"
	"github.com/pterm/pterm"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func NewMetricsClientSet(config *rest.Config) (*metricsv.Clientset, error) {
	return metricsv.NewForConfig(config)
}

func GetPodMetrics(clientset *metricsv.Clientset, namespace string) (*metricsv1beta1.PodMetricsList, error) {
	return clientset.MetricsV1beta1().PodMetricses(namespace).List(context.Background(), v1.ListOptions{})
}

func PrepareBarData(podMetrics *metricsv1beta1.PodMetricsList) []pterm.Bar {
	bars := []pterm.Bar{}
	for _, podMetric := range podMetrics.Items {
		for _, container := range podMetric.Containers {
			memoryUsage := container.Usage["memory"]
			bars = append(bars, pterm.Bar{Label: container.Name, Value: int(memoryUsage.Value())})
		}
	}
	return bars
}

func GetMemoryUsageBars(clientset *kubernetes.Clientset, metricsClientset *metricsv.Clientset, namespace string) ([]pterm.Bar, error) {
	var namespaces *corev1.NamespaceList
	var err error
	if namespace == "" {
		namespaces, err = kube.GetNamespaces(clientset)
		if err != nil {
			return nil, err
		}
	} else {
		namespaces = &corev1.NamespaceList{
			Items: []corev1.Namespace{
				{
					ObjectMeta: v1.ObjectMeta{
						Name: namespace,
					},
				},
			},
		}
	}

	bars := []pterm.Bar{}
	for _, ns := range namespaces.Items {
		podMetrics, err := GetPodMetrics(metricsClientset, ns.Name)
		if err != nil {
			return nil, err
		}
		bars = append(bars, PrepareBarData(podMetrics)...)
	}
	return bars, nil
}
