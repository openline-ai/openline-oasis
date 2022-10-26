#! /bin/bash

# Deploy Images
NAMESPACE_NAME="openline-development"
echo "script is $0"
OASIS_HOME="$(dirname $(readlink -f $0))/../../../"
echo "OASIS_HOME=$OASIS_HOME"

if [[ $(kubectl get namespaces) == *"$NAMESPACE_NAME"* ]];
  then
    echo " --- Continue deploy on namespace openline-development --- "
  else
    echo " --- Creating Openline Development namespace in minikube ---"
    kubectl create -f "$OPENLINE_HOME/deployment/k8s/openline-namespace.json"
    wait
fi

## Build Images
cd  $OASIS_HOME
docker build -t message-store -f message-store/Dockerfile .
docker build -t oasis-api -f oasis-api/Dockerfile .
docker build -t channels-api -f channels-api/Dockerfile .
cd oasis-voice/kamailio/;./build-docker.sh;cd $OASIS_HOME
cd oasis-voice/asterisk/;./build-docker.sh; cd $OASIS_HOME

minikube image load message-store:latest
minikube image load oasis-api:latest
minikube image load channels-api:latest
minikube image load ghcr.io/openline-ai/openline-oasis/openline-kamailio-server:latest
minikube image load ghcr.io/openline-ai/openline-oasis/openline-asterisk-server:latest

cd $OASIS_HOME/deployment/k8s/local-minikube
kubectl apply -f postgres/postgresql-configmap.yaml --namespace $NAMESPACE_NAME
kubectl apply -f postgres/postgresql-storage.yaml --namespace $NAMESPACE_NAME
kubectl apply -f postgres/postgresql-deployment.yaml --namespace $NAMESPACE_NAME
kubectl apply -f postgres/postgresql-service.yaml --namespace $NAMESPACE_NAME

kubectl apply -f apps-config/message-store.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/message-store-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-api.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-api-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/asterisk.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/kamailio.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/kamailio-k8s-service.yaml --namespace $NAMESPACE_NAME

cd $OASIS_HOME/oasis-voice/kamailio/sql
./build_db.sh local-kube
cd $OASIS_HOME/deployment/k8s/local-minikube



echo "run the following port forwarding commands"
echo kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 8080:8080
echo kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 5060:5060

