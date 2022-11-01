#!/bin/bash
docker buildx build -t ghcr.io/openline-ai/openline-oasis/openline-kamailio-server --platform linux/amd64 .
