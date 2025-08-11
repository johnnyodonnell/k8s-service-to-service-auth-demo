#!/bin/bash


# Clean up previous deployments
kubectl delete deployment demo-service-a
kubectl delete service demo-service-a
kubectl delete deployment demo-service-b
kubectl delete service demo-service-b

# Deploy services
kubectl create deployment demo-service-a --image=demo-service:current
kubectl set env deployment/demo-service-a SERVICE=http://demo-service-b:8081
kubectl expose deployment demo-service-a --type=ClusterIP --port=8080 --target-port=8080

kubectl create deployment demo-service-b --image=demo-service:current
kubectl set env deployment/demo-service-b SERVICE=http://demo-service-a:8080
kubectl expose deployment demo-service-b --type=ClusterIP --port=8081 --target-port=8080

