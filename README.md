# Kubectl Plugin: Display Kubernetes Memory Usage Bar Chart

This program fetches the memory usage of all pods in all namespaces from a Kubernetes cluster and displays it in a bar chart.

## Prerequisites

- Go 1.16 or later
- Access to a Kubernetes cluster
- Kubeconfig file (default looks at `$HOME/.kube/config`)

## Libraries Used

- [client-go](https://github.com/kubernetes/client-go) for interacting with the Kubernetes API
- [metrics-server](https://github.com/kubernetes-sigs/metrics-server) to fetch pod resource usage
- [pterm](https://github.com/pterm/pterm) to display the bar chart

## Installation

```shell
go build
cp kubectl-mem-usage /your/bin/path
```

## Usage

### Example: Display pods memory usage of specific namespace

<img src="https://raw.githubusercontent.com/jedipunkz/kubectl-mem-usage/main/pix/kubectl-mem-usage.png">


## Author

@jedipunkz

## License

Apache 2.0
