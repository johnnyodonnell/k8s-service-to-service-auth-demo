#!/bin/bash


docker build -t demo-service -f docker/Dockerfile .

docker tag demo-service:latest demo-service:current

minikube image load demo-service:current

