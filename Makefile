VERSION := 0.0.1

DOCKER_REG = damiannolan
DOCKER_IMAGE = eventing-init
DOCKER_IMAGE_TAG = $(VERSION)
USER = $(shell whoami)

docker-build:
	docker build -t $(DOCKER_REG)/$(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG) .

docker-build-dev:
	devdocker build -t $(DOCKER_REG)/$(DOCKER_IMAGE):$(USER) .

docker-push:
	docker push $(DOCKER_REG)/$(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG)

docker-push-dev:
	docker push  $(DOCKER_REG)/$(DOCKER_IMAGE):$(USER) 

.PHONY: \
	docker-build \
	docker-build-dev \
	docker-push \
	docker-push-dev 
