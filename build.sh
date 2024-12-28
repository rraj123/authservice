#!/bin/bash

source .env

# Build Docker image
docker image build ${BUILD_OPTS} -t ${REGISTRY}${IMAGE}${TAG} .

docker image push ${REGISTRY}${IMAGE}${TAG}