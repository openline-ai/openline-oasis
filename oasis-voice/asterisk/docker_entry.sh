#!/bin/sh

PUBLIC_IP=172.17.0.1 LOCAL_IP=172.17.0.2 /usr/sbin/asterisk_network_setup.sh
/usr/sbin/asterisk -T -W -U asterisk -p -vvvdddf