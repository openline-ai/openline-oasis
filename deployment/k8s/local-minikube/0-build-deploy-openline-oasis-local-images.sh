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
  cd $CUSTOMER_OS_HOME/deployment/scripts/old/
  ../0-get-config.sh
  cd "$OASIS_HOME/../"
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
    cd $CUSTOMER_OS_HOME/deployment/scripts/old/
    $CUSTOMER_OS_HOME/deployment/scripts/old/1-ubuntu-dependencies.sh
    if [ $INSTALLED_DOCKER == 1 ];
    then 
	    echo "Docker has just been installed"
	    echo "Please logout and log in for the group changes to take effect"
	    echo "Once logged back in, re-run this script to resume the installation"
	    exit
    fi
  fi
  if [ "x$(uname -s)" == "xDarwin" ]; 
  then
    echo "missing base dependencies, installing"
    cd $CUSTOMER_OS_HOME/deployment/scripts/old/
    $CUSTOMER_OS_HOME/deployment/scripts/old/1-mac-dependencies.sh
  fi
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
  cd $CUSTOMER_OS_HOME/deployment/scripts/old/
  $CUSTOMER_OS_HOME/deployment/scripts/old/2-base-install.sh
fi

if [ -z "$(kubectl get deployment customer-os-api -n $CUSTOMER_OS_NAME_SPACE)" ]; 
then
  echo "Installing Customer OS Aplicaitons"
  getCustomerOs
  cd $CUSTOMER_OS_HOME/deployment/scripts/old/
  $CUSTOMER_OS_HOME/deployment/scripts/old/3-deploy.sh
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
cd  $OASIS_HOME

if [ "x$1" == "xbuild" ]; then
  if [ "x$(lsb_release -i|cut -d: -f 2|xargs)" == "xUbuntu" ];
  then
    if [ -z "$(which protoc)" ]; 
    then
	    sudo apt-get update
	    sudo apt-get install -y unzip wget
	    cd /tmp/
	    wget https://github.com/protocolbuffers/protobuf/releases/download/v21.9/protoc-21.9-linux-x86_64.zip
	    unzip protoc-21.9-linux-x86_64.zip
	    sudo mv bin/protoc /usr/local/bin
	    sudo mv include/* /usr/local/include/
    fi
    if [ -z "$(which go)" ]; 
    then
	    sudo apt-get update
	    sudo apt-get install -y golang-go
	    mkdir -p ~/go/{bin,src,pkg}
	    export GOPATH="$HOME/go"
	    export GOBIN="$GOPATH/bin"
    fi
    if [ -z "$(which make)" ]; 
    then
	    sudo apt-get install make
    fi
  fi
  if [ "x$(uname -s)" == "xDarwin" ]; 
  then
	  brew install protobuf
  fi
  cd $OASIS_HOME/packages/server/channels-api;make install;make generate;cd $OASIS_HOME
  cd $OASIS_HOME/packages/server/oasis-api;make install;make generate;cd $OASIS_HOME

  docker build -t ghcr.io/openline-ai/openline-oasis/oasis-api:otter -f packages/server/oasis-api/Dockerfile ./packages/server/.
  docker build -t ghcr.io/openline-ai/openline-oasis/channels-api:otter -f packages/server/channels-api/Dockerfile ./packages/server/.
  docker build -t ghcr.io/openline-ai/openline-oasis/oasis-frontend-dev:otter --platform linux/amd64 --build-arg NODE_ENV=dev ./packages/apps/oasis/oasis-frontend
else
  docker pull ghcr.io/openline-ai/openline-oasis/oasis-api:otter
  docker pull ghcr.io/openline-ai/openline-oasis/channels-api:otter
  docker pull ghcr.io/openline-ai/openline-oasis/oasis-frontend-dev:otter
fi

minikube image load ghcr.io/openline-ai/openline-oasis/oasis-api:otter --daemon
minikube image load ghcr.io/openline-ai/openline-oasis/channels-api:otter --daemon
minikube image load ghcr.io/openline-ai/openline-oasis/oasis-frontend-dev:otter --daemon

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

kubectl rollout restart -n $NAMESPACE_NAME deployment/oasis-api
kubectl rollout restart -n $NAMESPACE_NAME deployment/channels-api

cd $OASIS_HOME/deployment/k8s/local-minikube
