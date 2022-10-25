#!/bin/bash
docker build -t ghcr.io/openline-ai/openline-oasis/openline-kamailio-server --build-arg ARCH=amd64/ .
