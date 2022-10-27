#!/bin/sh

docker run -d  --env-file .env.dev --name openline-asterisk-server ghcr.io/openline-ai/openline-oasis/openline-asterisk-server
