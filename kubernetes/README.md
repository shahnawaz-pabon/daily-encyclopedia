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
    - [Interact with the cluster](#interact-with-the-cluster)
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

#### Interact with the cluster

```sh
# Start a new terminal, and leave this running.
minikube dashboard
```

Next, return to the terminal where you initiated the `minikube start` command.

### kubectl

For installation instructions, please refer to the [Install tools](https://kubernetes.io/docs/tasks/tools/#kubectl) section (Manual Installation). Alternatively, Minikube can automatically download the necessary version of kubectl, allowing you to use it as follows:

```sh
minikube kubectl -- get pods -A
```

Simplify your workflow by including the following in your shell configuration:

```sh
alias kubectl="minikube kubectl --"
```
