#!/bin/bash


docker build -t kingod180/demo-service -f docker/Dockerfile .

docker push kingod180/demo-service

