#!/bin/sh

sed "s/#MY_PUBLIC_IP_ADDR#/$PUBLIC_IP/g" /etc/kamailio/network.conf.template| sed "s/#MY_PRIVATE_IP_ADDR#/$LOCAL_IP/g" > /etc/kamailio/network.conf

if [ -n "$ASTERISK_HOST" ]
then
    echo "0 sip:$ASTERISK_HOST 0 0" >> /etc/kamailio/dispatcher.list
fi
