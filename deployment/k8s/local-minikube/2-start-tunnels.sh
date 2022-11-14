#!/bin/bash

export NAMESPACE_NAME=oasis-dev
minikube tunnel --bind-address 127.0.0.1 &
if [ $(uname -m) == "x86_64" ];
then
  kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 8080:8080 &
fi

wait
