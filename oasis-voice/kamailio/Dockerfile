FROM kamailio/kamailio:5.5.5-bullseye

RUN sh -c 'apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -qq --assume-yes libpq5 python3-psycopg2'
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

RUN sh -c 'rm -rf /var/lib/apt/lists/*'
ENTRYPOINT ["/docker_entry.sh"]
