IMAGE_NAME = manganellidev/kyma-demo
TAG = latest
DEPLOYMENT_FILE = k8s/deployment.yaml

build-image:
	docker build -t $(IMAGE_NAME):$(TAG) .

push-image:
	docker push $(IMAGE_NAME):$(TAG)

.PHONY: build-image push-image deploy