package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/jedipunkz/kubectl-mem-usage/internal/kube"
	"github.com/jedipunkz/kubectl-mem-usage/internal/metrics"
	"github.com/pterm/pterm"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	optNamespace := flag.String("n", "", "namespace")
	flag.Parse()

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}

	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Println(err)
	}

	clientset, err := kube.NewClientSet(config)
	if err != nil {
		log.Println(err)
	}

	metricsClientset, err := metrics.NewMetricsClientSet(config)
	if err != nil {
		log.Println(err)
	}

	bars, err := metrics.GetMemoryUsageBars(clientset, metricsClientset, *optNamespace)
	if err != nil {
		log.Println(err)
	}

	pterm.DefaultBarChart.WithBars(bars).WithHorizontal().WithShowValue().WithWidth(30).Render()
}
