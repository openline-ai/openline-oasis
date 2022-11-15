#!/bin/bash

export NAMESPACE_NAME=oasis-dev
minikube tunnel --bind-address 127.0.0.1 &
wait
