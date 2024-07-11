# k8s-demo

## Prerequisites

- [Go](https://go.dev/doc/install)
- [docker desktop](https://www.docker.com/products/docker-desktop/)
- [docker hub](https://hub.docker.com/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start)
- [BTP trial account](https://cockpit.hanatrial.ondemand.com/)
- [VS Code](https://code.visualstudio.com/download) or any other

## Demo

### Build and push image

1. Build image

```bash
make build-image REPO_NAME=xpto

# Optional
docker run -p 8080:8080 xpto/xpto:latest
```

2. Login to docker hub

```bash
docker login
```

2. Push image

```bash
make push-image REPO_NAME=xpto
```

### Setup Environment

1. Create cluster

```bash
# Only on local environment
kind create cluster --config=kind/kind.yaml --name=kyma-demo

# Verify current context
kubectl config current-context

# Optional: Set target context
kubectl config use-context xpto
```

2. Create namespace

```bash
kubectl create namespace dev


# Create context for the namespace
kubectl config view
kubectl config set-context xpto-dev --namespace=dev --cluster=xpto --user=xpto

# Verify current context
kubectl config current-context

# Optional: Set target context
kubectl config use-context xpto
```

### Deployment

1. Adjust the image value with your repository name
2. Apply deployment

```bash
kubectl apply -f k8s/deployment.yaml

# Optional: Validate it is up and running
kubectl get pods

# Optional: Config port-forward to test it
kubectl port-forward pods/goserver-xpto 8080:8080
```

### Service

1. Apply service

```bash
kubectl apply -f k8s/service.yaml

# Optional: Validate it is up and running
kubectl get services

# Optional: Config port-forward to test it
kubectl port-forward pods/goserver-xpto 8080:80
```

### ConfigMap

1. Apply configmap

```bash
kubectl apply -f k8s/configmap.yaml

# Optional: Validate it is up and running
kubectl get configmaps

# Refresh deployment with updated config values
kubectl rollout restart deployment goserver
```
