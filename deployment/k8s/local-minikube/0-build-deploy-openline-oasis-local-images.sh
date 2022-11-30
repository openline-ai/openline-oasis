#! /bin/bash

# Deploy Images
NAMESPACE_NAME="oasis-dev"
CUSTOMER_OS_NAME_SPACE="openline"
echo "script is $0"
OASIS_HOME="$(dirname $(readlink -f $0))/../../../"
echo "OASIS_HOME=$OASIS_HOME"
CUSTOMER_OS_HOME="$OASIS_HOME/../openline-customer-os"
if [[ $(kubectl get namespaces) == *"$NAMESPACE_NAME"* ]];
  then
    echo " --- Continue deploy on namespace $NAMESPACE_NAME --- "
  else
    echo " --- Creating $NAMESPACE_NAME namespace in minikube ---"
    kubectl create -f "$OASIS_HOME/deployment/k8s/local-minikube/oasis-dev.json"
    wait
fi

## Build Images
cd  $OASIS_HOME

  docker build -t ghcr.io/openline-ai/openline-oasis/oasis-api:latest -f packages/server/oasis-api/Dockerfile ./packages/server/.
  docker build -t ghcr.io/openline-ai/openline-oasis/channels-api:latest -f packages/server/channels-api/Dockerfile ./packages/server/.

cd $OASIS_HOME/deployment/k8s/local-minikube

kubectl apply -f apps-config/oasis-api.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-api-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-api-k8s-loadbalancer-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api-k8s-loadbalancer-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-frontend.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-frontend-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-frontend-k8s-loadbalancer-service.yaml --namespace $NAMESPACE_NAME


cd $OASIS_HOME/deployment/k8s/local-minikube
