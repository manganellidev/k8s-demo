# kyma-demo

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
make build-image
```

2. Login to docker hub

```bash
docker login
```

2. Push image

```bash
make push-image
```

### Setup Environment

1. Create cluster

```bash
# Only using Kind locally
kind create cluster --config=kind/kind.yaml --name=kyma-demo
```

2. Create namespace

```bash
kubectl create namespace kyma-demo
```

### Deployment

1. Apply deployment yaml

```bash
kubectl apply -f k8s/deployment.yaml
```

### Service

1. Apply service yaml

```bash
kubectl apply -f k8s/service.yaml
```
