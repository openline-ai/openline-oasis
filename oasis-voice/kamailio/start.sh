#!/bin/sh

docker run -d -p 8080:8080 --env-file .env.dev --name openline-kamailio-server openline.ai/openline-kamailio-server
