From ubuntu:jammy

ARG opus_codec=asterisk-18.0/x86-64/codec_opus-18.0_current-x86_64
RUN sh -c 'apt-get update && DEBIAN_FRONTEND=noninteractive apt install -y software-properties-common'
RUN sh -c 'add-apt-repository universe && apt-get update'
RUN sh -c 'DEBIAN_FRONTEND=noninteractive apt-get install -y asterisk sox curl'
RUN sh -c 'mkdir -p /usr/src/codecs/opus'
RUN sh -c 'cd /usr/src/codecs/opus && curl -sL http://downloads.digium.com/pub/telephony/codec_opus/${opus_codec}.tar.gz | tar --strip-components 1 -xz'
RUN sh -c 'cp /usr/src/codecs/opus/*.so /usr/lib/x86_64-linux-gnu/asterisk/modules/'
RUN sh -c 'cp /usr/src/codecs/opus/codec_opus_config-en_US.xml /usr/share/asterisk/documentation/'
RUN sh -c 'rm /usr/lib/x86_64-linux-gnu/asterisk/modules/format_ogg_opus_open_source.so'

COPY docker_entry.sh /
COPY scripts/asterisk_network_setup.sh /usr/sbin/
COPY conf/extensions.conf /etc/asterisk/
COPY conf/pjsip.conf.template /etc/asterisk/
COPY conf/rtp.conf.template /etc/asterisk/

RUN sh -c 'rm -rf /var/lib/apt/lists/*'

EXPOSE 10000-20000/udp

ENTRYPOINT ["/docker_entry.sh"]