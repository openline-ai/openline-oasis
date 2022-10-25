#!/bin/sh

/etc/kamailio/genconf.sh
PUBLIC_IP=172.17.0.1 LOCAL_IP=172.17.0.2 /usr/sbin/kamailio_network_setup.sh
touch /etc/kamailio/dispatcher.list
kamailio -DD -E
