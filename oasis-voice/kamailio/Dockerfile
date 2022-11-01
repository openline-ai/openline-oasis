FROM amd64/debian:bullseye

ENV DEBIAN_FRONTEND=noninteractive
RUN rm -rf /var/lib/apt/lists/* && apt-get update && apt-get install -qq --assume-yes gnupg wget
# kamailio repo
RUN echo "deb http://deb.kamailio.org/kamailio55 bullseye main" >   /etc/apt/sources.list.d/kamailio.list
RUN wget -O- http://deb.kamailio.org/kamailiodebkey.gpg | apt-key add -
RUN apt-get update && apt-get install -qq --no-install-recommends --no-install-suggests kamailio=5.5.5+bpo11 kamailio-autheph-modules=5.5.5+bpo11 kamailio-berkeley-bin=5.5.5+bpo11 kamailio-cnxcc-modules=5.5.5+bpo11 kamailio-cpl-modules=5.5.5+bpo11 kamailio-extra-modules=5.5.5+bpo11 kamailio-json-modules=5.5.5+bpo11 kamailio-lwsc-modules=5.5.5+bpo11 kamailio-memcached-modules=5.5.5+bpo11 kamailio-mqtt-modules=5.5.5+bpo11 kamailio-outbound-modules=5.5.5+bpo11 kamailio-phonenum-modules=5.5.5+bpo11 kamailio-postgres-modules=5.5.5+bpo11 kamailio-presence-modules=5.5.5+bpo11 kamailio-python3-modules=5.5.5+bpo11 kamailio-rabbitmq-modules=5.5.5+bpo11 kamailio-redis-modules=5.5.5+bpo11 kamailio-sctp-modules=5.5.5+bpo11 kamailio-secsipid-modules=5.5.5+bpo11 kamailio-snmpstats-modules=5.5.5+bpo11 kamailio-systemd-modules=5.5.5+bpo11 kamailio-tls-modules=5.5.5+bpo11 kamailio-utils-modules=5.5.5+bpo11 kamailio-websocket-modules=5.5.5+bpo11 kamailio-xml-modules=5.5.5+bpo11 libpq5 python3-psycopg2 && apt-get clean && rm -rf /var/lib/apt/lists/*

VOLUME /etc/kamailio

#Copy the config file onto the Filesystem of the Docker instance
COPY conf/kamailio.cfg /etc/kamailio/
COPY conf/kamailio.py /etc/kamailio/
COPY conf/KamailioDatabase.py /etc/kamailio/
COPY conf/network.conf.template /etc/kamailio/
COPY docker_entry.sh /
COPY scripts/genconf.sh /etc/kamailio/
COPY scripts/kamailio_network_setup.sh /usr/sbin/


#Expose port 5060 (SIP) for TCP and UDP
EXPOSE 5060
EXPOSE 5060/udp

EXPOSE 8080
EXPOSE 4443

RUN apt-get clean && rm -rf /var/lib/apt/lists/*
ENTRYPOINT ["/docker_entry.sh"]
