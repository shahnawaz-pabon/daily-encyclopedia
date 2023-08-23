<div align="center">
  <a href="https://kubernetes.io/">
    <img alt="kubernetes" src="../logos/kubernetes.png" height="100" width="100"/>
  </a>
  <h1>Kubernetes</h1>
</div>

# Table of Contents

- [Prerequisites](#prerequisites)
  - [minikube](#minikube)
  - [kubectl](#kubectl)

## Prerequisites

### minikube

`minikube` is local Kubernetes, focusing on making it easy to learn and develop for Kubernetes.
All you need is Docker (or similarly compatible) container or a Virtual Machine environment, and Kubernetes is a single command away: `minikube start`

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
