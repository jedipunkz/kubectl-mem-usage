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

### Display pods memory usage of all namespaces

```shell
$ kubectl mem usage

argo-rollouts                      ██                              71667712
nginx                                                              6303744
istio-proxy                        █                               40837120
istio-proxy                        █                               40820736
nginx                                                              11440128
istio-proxy                        █                               41873408
nginx                                                              12877824
istio-operator                     ████                            136933376
istio-proxy                        █                               39366656
istio-proxy                        ████                            115707904
discovery                          ████                            118132736
prometheus-server                  ██████████████████████████████  829476864
prometheus-server-configmap-reload                                 9961472
coredns                            ██                              58937344
etcd                               ██████████████                  388603904
kube-apiserver                     ██████████████████████          612376576
kube-controller-manager            █████                           138285056
kube-proxy                         █                               54198272
kube-scheduler                     ██                              59260928
metrics-server                                                     22220800
storage-provisioner                                                13029376
```

### Display pods memory usage of specific namespace

```shell
$ kubectl mem usage -n kube-system

coredns                 ██                              58941440
etcd                    ███████████████████             389341184
kube-apiserver          ██████████████████████████████  609923072
kube-controller-manager ██████                          131575808
kube-proxy              ██                              54202368
kube-scheduler          ██                              59265024
metrics-server          █                               22208512
storage-provisioner                                     12922880
```

## Author

@jedipunkz

## License

Apache 2.0
