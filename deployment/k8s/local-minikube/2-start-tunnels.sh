#!/bin/bash

export NAMESPACE_NAME=oasis-dev
kubectl port-forward --namespace $NAMESPACE_NAME svc/oasis-frontend-service 3006:3006 &
kubectl port-forward --namespace $NAMESPACE_NAME svc/oasis-api-service 8006:8006 &
kubectl port-forward --namespace $NAMESPACE_NAME svc/channels-api-service 8013:8013 &

if [ $(uname -m) == "x86_64" ];
then
  kubectl port-forward --namespace $NAMESPACE_NAME svc/kamailio-service 8080:8080 &
fi

wait
