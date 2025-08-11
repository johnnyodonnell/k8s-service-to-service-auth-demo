#!/bin/bash


kubectl delete deployment demo-service-a
kubectl delete service demo-service-a
kubectl delete deployment demo-service-b
kubectl delete service demo-service-b

