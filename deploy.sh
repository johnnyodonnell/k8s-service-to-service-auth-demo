#!/bin/bash


./cleanup.sh

# Deploy services
kubectl apply -f kubernetes/services.yaml

