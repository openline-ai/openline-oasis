#!/bin/sh

echo "#!substdef \"!EPHEMERAL_AUTH_SECRET!$AUTH_SECRET!g\"" > /etc/kamailio/local.conf
echo "#!substdef \"!DBURL!postgres://$SQL_USER:$SQL_PASSWORD@$SQL_HOST/$SQL_DATABASE!g\"" >> /etc/kamailio/local.conf

echo "[database]" > /etc/kamailio/config.ini
echo "host = $SQL_HOST" >> /etc/kamailio/config.ini
echo "database = $SQL_DATABASE" >> /etc/kamailio/config.ini
echo "user = $SQL_USER" >> /etc/kamailio/config.ini
echo "password = $SQL_PASSWORD" >> /etc/kamailio/config.ini



