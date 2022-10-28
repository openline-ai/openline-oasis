#!/bin/bash

 docker run -d --network=host coturn/coturn -v -z -n -X $(hostname -I|cut -f1 -d ' ') -L 0.0.0.0 --min-port=10000 --max-port=20000 

 echo "be sure port 3478 is tunneled"