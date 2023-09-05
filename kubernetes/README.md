<div align="center">
  <a href="https://kubernetes.io/">
    <img alt="kubernetes" src="../logos/kubernetes.png" height="100" width="100"/>
  </a>
  <h1>Kubernetes</h1>
</div>

# Table of Contents

- [Prerequisites](#prerequisites)
  - [minikube](#minikube)
    - [Linux Installation](#linux-installation)
    - [Start the cluster](#start-the-cluster)
  - [kubectl](#kubectl)

## Prerequisites

### minikube

`minikube` simplifies Kubernetes for local development, providing an accessible learning and development platform. With just Docker (or a compatible container system) or a Virtual Machine setup, you can start Kubernetes with a single command: `minikube start`.

#### Linux Installation

```sh
# To install the latest minikube stable release on x86-64 Linux using Debian package
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb
```

#### Start the cluster

In case `minikube` encounters issues while starting, refer to the [drivers page](https://minikube.sigs.k8s.io/docs/drivers/) for assistance in configuring a compatible container or virtual machine manager.

### kubectl

[Install tools](https://kubernetes.io/docs/tasks/tools/#kubectl) for installation instructions(Manual Installation).
Alternatively, minikube can download the appropriate version of kubectl and you should be able to use it like this:

```sh
minikube kubectl -- get pods -A
```

You can also make your life easier by adding the following to your shell config:

```sh
alias kubectl="minikube kubectl --"
```
