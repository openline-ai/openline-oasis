#! /bin/sh

# Deploy Images
NAMESPACE_NAME="openline-development"

## Build Images
cd  $OASIS_HOME
docker build -t message-store -f message-store/Dockerfile .
docker build -t oasis-api -f oasis-api/Dockerfile .
docker build -t channels-api -f channels-api/Dockerfile .

minikube image load message-store:latest
minikube image load oasis-api:latest
minikube image load channels-api:latest

cd $OASIS_HOME/deployment/k8s/local-minikube
kubectl create namespace $NAMESPACE_NAME
kubectl apply -f apps-config/message-store.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/message-store-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-api.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-api-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api-k8s-service.yaml --namespace $NAMESPACE_NAME
