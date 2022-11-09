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
  if [ $INSTALLED_DOCKER == 1 ];
  then 
    echo "Docker has just been installed"
    echo "Please logout and log in for the group changes to take effect"
    echo "Once logged back in, re-run this script to resume the installation"
    exit
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
  cd $OASIS_HOME/message-store;make install;make generate;cd $OASIS_HOME
  cd $OASIS_HOME/channels-api;make install;make generate;cd $OASIS_HOME

  docker build -t ghcr.io/openline-ai/openline-oasis/message-store:otter -f message-store/Dockerfile .
  docker build -t ghcr.io/openline-ai/openline-oasis/oasis-api:otter -f oasis-api/Dockerfile .
  docker build -t ghcr.io/openline-ai/openline-oasis/channels-api:otter -f channels-api/Dockerfile .
  docker build -t ghcr.io/openline-ai/openline-oasis/oasis-frontend-dev:otter --platform linux/amd64 --build-arg NODE_ENV=dev oasis-frontend
  if [ $(uname -m) == "x86_64" ];
  then
    cd oasis-voice/kamailio/;docker build -t ghcr.io/openline-ai/openline-oasis/openline-kamailio-server:otter .;cd $OASIS_HOME
    cd oasis-voice/asterisk/;docker build -t ghcr.io/openline-ai/openline-oasis/openline-asterisk-server:otter .;cd $OASIS_HOME
  fi
else
  docker pull ghcr.io/openline-ai/openline-oasis/message-store:otter
  docker pull ghcr.io/openline-ai/openline-oasis/oasis-api:otter
  docker pull ghcr.io/openline-ai/openline-oasis/channels-api:otter
  docker pull ghcr.io/openline-ai/openline-oasis/oasis-frontend-dev:otter
  if [ $(uname -m) == "x86_64" ];
  then
    docker pull ghcr.io/openline-ai/openline-oasis/openline-kamailio-server:otter
    docker pull ghcr.io/openline-ai/openline-oasis/openline-asterisk-server:otter
  fi


fi

minikube image load ghcr.io/openline-ai/openline-oasis/message-store:otter --daemon
minikube image load ghcr.io/openline-ai/openline-oasis/oasis-api:otter --daemon
minikube image load ghcr.io/openline-ai/openline-oasis/channels-api:otter --daemon
minikube image load ghcr.io/openline-ai/openline-oasis/oasis-frontend-dev:otter --daemon

if [ $(uname -m) == "x86_64" ];
then
  minikube image load ghcr.io/openline-ai/openline-oasis/openline-kamailio-server:otter --daemon
  minikube image load ghcr.io/openline-ai/openline-oasis/openline-asterisk-server:otter --daemon
fi

if [ $(uname -m) == "x86_64" ];
then
  cd $OASIS_HOME/oasis-voice/kamailio/sql
  SQL_USER=openline-oasis SQL_DATABABASE=openline-oasis ./build_db.sh local-kube
fi
  
cd $OASIS_HOME/deployment/k8s/local-minikube

kubectl apply -f apps-config/message-store.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/message-store-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-api.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-api-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/channels-api-k8s-service.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-frontend.yaml --namespace $NAMESPACE_NAME
kubectl apply -f apps-config/oasis-frontend-k8s-service.yaml --namespace $NAMESPACE_NAME

if [ $(uname -m) == "x86_64" ];
then
  kubectl apply -f apps-config/asterisk.yaml --namespace $NAMESPACE_NAME
  kubectl apply -f apps-config/asterisk-k8s-service.yaml --namespace $NAMESPACE_NAME
  kubectl apply -f apps-config/kamailio.yaml --namespace $NAMESPACE_NAME
  kubectl apply -f apps-config/kamailio-k8s-service.yaml --namespace $NAMESPACE_NAME
fi

kubectl rollout restart -n $NAMESPACE_NAME deployment/message-store
kubectl rollout restart -n $NAMESPACE_NAME deployment/oasis-api
kubectl rollout restart -n $NAMESPACE_NAME deployment/channels-api



cd $OASIS_HOME/message-store/sql
SQL_USER=openline-oasis SQL_DATABABASE=openline-oasis ./build_db.sh local-kube
  
cd $OASIS_HOME/deployment/k8s/local-minikube
echo "run the following port forwarding commands"
echo kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 8080:8080
echo kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 5060:5060
