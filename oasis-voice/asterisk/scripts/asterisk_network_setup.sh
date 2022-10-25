#!/bin/sh

sed "s/!PUBLIC_IP!/$PUBLIC_IP/g" /etc/asterisk/pjsip.conf.template| sed "s/!LOCAL_IP!/$LOCAL_IP/g" > /etc/asterisk/pjsip.conf

sed "s/!PUBLIC_IP!/$PUBLIC_IP/g" /etc/asterisk/rtp.conf.template| sed "s/!LOCAL_IP!/$LOCAL_IP/g" > /etc/asterisk/rtp.conf
