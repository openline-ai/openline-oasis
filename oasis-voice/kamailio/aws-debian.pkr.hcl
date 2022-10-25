variable "environment" {
	type=string
	default="uat-ninja"
	sensitive=false
}
data "amazon-parameterstore" "auth_secret" {
  name = "/config/kamailio-server_${var.environment}/auth_secret"
  with_decryption = false
}

data "amazon-parameterstore" "db_host" {
  name = "/config/kamailio-server_${var.environment}/db_host"
  with_decryption = false
}

data "amazon-parameterstore" "db_user" {
  name = "/config/kamailio-server_${var.environment}/db_user"
  with_decryption = false
}

data "amazon-parameterstore" "db_database" {
  name = "/config/kamailio-server_${var.environment}/db_database"
  with_decryption = false
}

data "amazon-parameterstore" "db_password" {
  name = "/config/kamailio-server_${var.environment}/db_password"
  with_decryption = true
}

# usage example of the data source output
locals {
  auth_secret   = data.amazon-parameterstore.auth_secret.value
  db_user   = data.amazon-parameterstore.db_user.value
  db_database   = data.amazon-parameterstore.db_database.value
  db_host   = data.amazon-parameterstore.db_host.value
  db_password   = data.amazon-parameterstore.db_password.value
}

packer {
  required_plugins {
    amazon = {
      version = ">= 0.0.2"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

source "amazon-ebs" "debian" {
  ami_name      = "openline-voice-kamailio_${var.environment}"
  instance_type = "t2.micro"
  region        = "eu-west-2"
  source_ami_filter {
    filters = {
      name                = "debian-11-amd64-*"
      root-device-type    = "ebs"
      virtualization-type = "hvm"
    }
    most_recent = true
    owners      = ["136693071363"]
  }
  ssh_username = "admin"
}

build {
  name = "openline-voice-kamailio-server"
  sources = [
    "source.amazon-ebs.debian"
  ]
  provisioner "shell" { 
	  inline=  [
	    "sudo sh -c 'rm -rf /var/lib/apt/lists/* && apt-get update &&   DEBIAN_FRONTEND=noninteractive apt-get install -qq --assume-yes gnupg wget'",
	    "sudo sh -c 'echo \"deb http://deb.kamailio.org/kamailio55 bullseye main\" >   /etc/apt/sources.list.d/kamailio.list'",
	    "sudo sh -c 'wget -O- http://deb.kamailio.org/kamailiodebkey.gpg | apt-key add -'",
	    "sudo sh -c 'apt-get update &&   DEBIAN_FRONTEND=noninteractive apt-get install -qq --assume-yes kamailio=5.5.5+bpo11 kamailio-autheph-modules=5.5.5+bpo11 kamailio-berkeley-bin=5.5.5+bpo11 kamailio-berkeley-modules=5.5.5+bpo11 kamailio-cnxcc-modules=5.5.5+bpo11 kamailio-cpl-modules=5.5.5+bpo11 kamailio-dbg=5.5.5+bpo11 kamailio-erlang-modules=5.5.5+bpo11 kamailio-extra-modules=5.5.5+bpo11 kamailio-geoip-modules=5.5.5+bpo11 kamailio-geoip2-modules=5.5.5+bpo11 kamailio-ims-modules=5.5.5+bpo11 kamailio-json-modules=5.5.5+bpo11 kamailio-kazoo-modules=5.5.5+bpo11 kamailio-ldap-modules=5.5.5+bpo11 kamailio-lua-modules=5.5.5+bpo11 kamailio-lwsc-modules=5.5.5+bpo11 kamailio-memcached-modules=5.5.5+bpo11 kamailio-mongodb-modules=5.5.5+bpo11 kamailio-mono-modules=5.5.5+bpo11 kamailio-mqtt-modules=5.5.5+bpo11 kamailio-mysql-modules=5.5.5+bpo11 kamailio-nth=5.5.5+bpo11 kamailio-outbound-modules=5.5.5+bpo11 kamailio-perl-modules=5.5.5+bpo11 kamailio-phonenum-modules=5.5.5+bpo11 kamailio-postgres-modules=5.5.5+bpo11 kamailio-presence-modules=5.5.5+bpo11 kamailio-python-modules=5.5.5+bpo11 kamailio-python3-modules=5.5.5+bpo11 kamailio-rabbitmq-modules=5.5.5+bpo11 kamailio-radius-modules=5.5.5+bpo11 kamailio-redis-modules=5.5.5+bpo11 kamailio-ruby-modules=5.5.5+bpo11 kamailio-sctp-modules=5.5.5+bpo11 kamailio-secsipid-modules=5.5.5+bpo11 kamailio-snmpstats-modules=5.5.5+bpo11 kamailio-sqlite-modules=5.5.5+bpo11 kamailio-systemd-modules=5.5.5+bpo11 kamailio-tls-modules=5.5.5+bpo11 kamailio-unixodbc-modules=5.5.5+bpo11 kamailio-utils-modules=5.5.5+bpo11 kamailio-websocket-modules=5.5.5+bpo11 kamailio-xml-modules=5.5.5+bpo11 kamailio-xmpp-modules=5.5.5+bpo11'",
	    "sudo sh -c 'DEBIAN_FRONTEND=noninteractive apt-get install -qq --assume-yes libpq5 python3-psycopg2'",
	    "mkdir /tmp/kamailio/",
	  ]
  }

  provisioner "file" {
    source = "conf"
    destination = "/tmp/kamailio/"
  }
  provisioner "file" {
    source = "scripts"
    destination = "/tmp/kamailio/"
  }
  provisioner "file" {
    source = "sql"
    destination = "/tmp/kamailio/"
  }
  provisioner "shell" {
    inline=  [
      "sudo sh -c 'mv /tmp/kamailio/conf/* /etc/kamailio/'",
      "sudo sh -c 'mv /tmp/kamailio/scripts/genconf.sh /etc/kamailio/'",
      "sudo sh -c 'mv /tmp/kamailio/scripts/kamailio_network_setup.sh /usr/sbin/'",
      "sudo sh -c 'mv /tmp/kamailio/scripts/kamailio.service /lib/systemd/system/'",
      "sudo sh -c 'chown kamailio:kamailio /etc/kamailio/'",
      "sudo sh -c 'AUTH_SECRET=${local.auth_secret} SQL_HOST=${local.db_host} SQL_USER=${local.db_user} SQL_PASSWORD=${local.db_password} SQL_DATABASE=${local.db_database} /etc/kamailio/genconf.sh'",
      "sudo sh -c 'touch /etc/kamailio/dispatcher.list'",
    ]
  }
}


