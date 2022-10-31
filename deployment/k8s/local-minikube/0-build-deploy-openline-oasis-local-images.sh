#! /bin/bash

# Deploy Images
NAMESPACE_NAME="oasis-dev"
CUSTOMER_OS_NAME_SPACE="openline"
echo "script is $0"
OASIS_HOME="$(dirname $(readlink -f $0))/../../../"
echo "OASIS_HOME=$OASIS_HOME"
CUSTOMER_OS_HOME="$OASIS_HOME/../openline-customer-os"
function getCustomerOs () {
  if [ ! -d $CUSTOMER_OS_HOME ];
  then
    cd "$OASIS_HOME/../"
    git clone https://github.com/openline-ai/openline-customer-os.git
  fi
}

if [ -z "$(which kubectl)" ] || [ -z "$(which docker)" ] || [ -z "$(which minikube)" ] ; 
then
  if [ -z "$(which docker)" ]; 
  then
    INSTALLED_DOCKER=1
  else
    INSTALLED_DOCKER=0
  fi
  getCustomerOs
  if [ "x$(lsb_release -i|cut -d: -f 2|xargs)" == "xUbuntu" ];
  then
    echo "missing base dependencies, installing"
    $CUSTOMER_OS_HOME/deployment/k8s/local-minikube/0-ubuntu-install-prerequisites.sh
  fi
  if [ "x$(uname -s)" == "xDarwin" ]; 
  then
    echo "Base env not ready, follow up the setup procedure at the following link"
    echo "https://github.com/openline-ai/openline-customer-os/tree/otter/deployment/k8s/local-minikube#setup-environment-for-osx"
    exit
  fi
  if [ INSTALLED_DOCKER == 1 ];
  then 
    echo "Docker has just been installed"
    echo "Please logout and log in for the group changes to take effect"
    echo "Once logged back in, re-run this script to resume the installation"
    exit
fi

MINIKUBE_STATUS=$(minikube status)
MINIKUBE_STARTED_STATUS_TEXT='Running'
if [[ "$MINIKUBE_STATUS" == *"$MINIKUBE_STARTED_STATUS_TEXT"* ]];
  then
     echo " --- Minikube already started --- "
  else
     eval $(minikube docker-env)
     minikube start &
     wait
fi

if [[ $(kubectl get namespaces) == *"$CUSTOMER_OS_NAME_SPACE"* ]];
then
  echo "Customer OS Base already installed"
else
  echo "Installing Customer OS Base"
  getCustomerOs
  $CUSTOMER_OS_HOME/deployment/k8s/local-minikube/1-deploy-customer-os-base-infrastructure-local.sh
fi

if [ -z "$(kubectl get deployment customer-os-api -n $CUSTOMER_OS_NAME_SPACE)" ]; 
then
  echo "Installing Customer OS Aplicaitons"
  getCustomerOs
  $CUSTOMER_OS_HOME/deployment/k8s/local-minikube/2-build-deploy-customer-os-local-images.sh
fi  

if [[ $(kubectl get namespaces) == *"$NAMESPACE_NAME"* ]];
  then
    echo " --- Continue deploy on namespace $NAMESPACE_NAME --- "
  else
    echo " --- Creating $NAMESPACE_NAME namespace in minikube ---"
    kubectl create -f "$OASIS_HOME/deployment/k8s/local-minikube/oasis-dev.json"
    wait
fi

## Build Images
cd $OASIS_HOME/deployment/k8s/local-minikube

minikube image load postgres:13.4 --pull

kubectl apply -f postgres/postgresql-configmap.yaml --namespace $NAMESPACE_NAME
kubectl apply -f postgres/postgresql-storage.yaml --namespace $NAMESPACE_NAME
kubectl apply -f postgres/postgresql-deployment.yaml --namespace $NAMESPACE_NAME
kubectl apply -f postgres/postgresql-service.yaml --namespace $NAMESPACE_NAME

cd  $OASIS_HOME

if [ "x$1" == "xbuild" ]; then
  minikube image build -t ghcr.io/openline-ai/openline-oasis/message-store:otter -f message-store/Dockerfile .
  minikube image build -t ghcr.io/openline-ai/openline-oasis/oasis-api:otter -f oasis-api/Dockerfile .
  minikube image build -t ghcr.io/openline-ai/openline-oasis/channels-api:otter -f channels-api/Dockerfile .
  cd oasis-voice/kamailio/;minikube image build -t ghcr.io/openline-ai/openline-oasis/openline-kamailio-server:otter .;cd $OASIS_HOME
  cd oasis-voice/asterisk/;minikube image build -t ghcr.io/openline-ai/openline-oasis/openline-asterisk-server:otter .;cd $OASIS_HOME
else
  docker pull ghcr.io/openline-ai/openline-oasis/message-store:otter
  docker pull ghcr.io/openline-ai/openline-oasis/oasis-api:otter
  docker pull ghcr.io/openline-ai/openline-oasis/channels-api:otter
  docker pull ghcr.io/openline-ai/openline-oasis/openline-kamailio-server:otter
  docker pull ghcr.io/openline-ai/openline-oasis/openline-asterisk-server:otter 

  minikube image load ghcr.io/openline-ai/openline-oasis/message-store:otter
  minikube image load ghcr.io/openline-ai/openline-oasis/oasis-api:otter
  minikube image load ghcr.io/openline-ai/openline-oasis/channels-api:otter
  minikube image load ghcr.io/openline-ai/openline-oasis/openline-kamailio-server:otter
  minikube image load ghcr.io/openline-ai/openline-oasis/openline-asterisk-server:otter 

fi


cd $OASIS_HOME/oasis-voice/kamailio/sql
SQL_USER=openline-oasis SQL_DATABABASE=openline-oasis ./build_db.sh local-kube
cd $OASIS_HOME/deployment/k8s/local-minikube

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


kubectl rollout restart -n $NAMESPACE_NAME deployment/message-store
kubectl rollout restart -n $NAMESPACE_NAME deployment/oasis-api
kubectl rollout restart -n $NAMESPACE_NAME deployment/channels-api

echo "run the following port forwarding commands"
echo kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 8080:8080
echo kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 5060:5060
