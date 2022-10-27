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

if [ "x$1" == "xbuild" ]; then
  minikube image build -t ghcr.io/openline-ai/openline-oasis/message-store -f message-store/Dockerfile .
  minikube image build -t ghcr.io/openline-ai/openline-oasis/oasis-api -f oasis-api/Dockerfile .
  minikube image build -t ghcr.io/openline-ai/openline-oasis/channels-api -f channels-api/Dockerfile .
  cd oasis-voice/kamailio/;minikube image build -t ghcr.io/openline-ai/openline-oasis/openline-kamailio-server .;cd $OASIS_HOME
  cd oasis-voice/asterisk/;minikube image build -t ghcr.io/openline-ai/openline-oasis/openline-asterisk-server .;cd $OASIS_HOME
else
  minikube pull ghcr.io/openline-ai/openline-oasis/message-store
  minikube pull ghcr.io/openline-ai/openline-oasis/oasis-api
  minikube pull ghcr.io/openline-ai/openline-oasis/channels-api
  minikube pull ghcr.io/openline-ai/openline-oasis/openline-kamailio-server
  minikube pull ghcr.io/openline-ai/openline-oasis/openline-asterisk-server

fi

minikube image pull postgres:13.4

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
kubectl apply -f apps-config/asterisk-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/kamailio.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/kamailio-k8s-service.yaml --namespace $NAMESPACE_NAME

cd $OASIS_HOME/oasis-voice/kamailio/sql
SQL_USER=openline-oasis SQL_DATABABASE=openline-oasis ./build_db.sh local-kube
cd $OASIS_HOME/deployment/k8s/local-minikube



echo "run the following port forwarding commands"
echo kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 8080:8080
echo kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 5060:5060

