# Concepts

## Kubernetes (k8s)

> Kubernetes is an open source container orchestration engine for automating deployment, scaling, and management of containerized applications.

Source: [k8s home](https://kubernetes.io/docs/home/)

#### Main points:

- K8s works through a set of APIs
- kubectl (kube control) is the CLI to interact with these APIs
- Everything is based on states
- K8s use objects (yaml files) to represent the states

### Cluster / Node

> A Kubernetes cluster consists of a set of worker machines, called nodes, that run containerized applications. Every cluster has at least one worker node.

Source: [k8s Components](https://kubernetes.io/docs/concepts/overview/components/)

### Kubernetes Objects

#### Pods

> Pods are the smallest deployable units of computing that you can create and manage in Kubernetes.

Source: [k8s Pods](https://kubernetes.io/docs/concepts/workloads/pods/)

#### ReplicaSet

> A ReplicaSet's purpose is to maintain a stable set of replica Pods running at any given time. As such, it is often used to guarantee the availability of a specified number of identical Pods.

Source: [k8s ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/)

#### Deployment

> A Deployment provides declarative updates for Pods and ReplicaSets.

> You describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.

Source: [k8s Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

#### Service

> In Kubernetes, a Service is a method for exposing a network application that is running as one or more Pods in your cluster.

Source: [k8s Service](https://kubernetes.io/docs/concepts/services-networking/service/)

#### ConfigMaps

> A ConfigMap is an API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume.

> A ConfigMap allows you to decouple environment-specific configuration from your container images, so that your applications are easily portable.

Source: [k8s ConfigMap](https://kubernetes.io/docs/concepts/configuration/configmap/)

#### HorizontalPodAutoscaler

> A HorizontalPodAutoscaler (HPA for short) automatically updates a workload resource (such as a Deployment or StatefulSet), with the aim of automatically scaling the workload to match demand.

Source: [k8s HorizontalPodAutoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/)
