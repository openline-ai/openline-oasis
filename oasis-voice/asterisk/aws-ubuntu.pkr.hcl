packer {
  required_plugins {
    amazon = {
      version = ">= 0.0.2"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

locals {
  opus_codec  = "asterisk-18.0/x86-64/codec_opus-18.0_current-x86_64"
}

source "amazon-ebs" "ubuntu" {
  ami_name      = "asterisk-server-ami"
  instance_type = "t2.micro"
  region        = "eu-west-2"
  source_ami_filter {
    filters = {
      name                = "ubuntu/images/*ubuntu-jammy-22.04-amd64-server-*"
      root-device-type    = "ebs"
      virtualization-type = "hvm"
    }
    most_recent = true
    owners      = ["099720109477"]
  }
  ssh_username = "ubuntu"
}

build {
  name    = "build-asterisk-image"
  sources = [
    "source.amazon-ebs.ubuntu"
  ]
  provisioner "shell" {
    inline = [
      "sudo sh -c 'add-apt-repository universe && apt-get update'",
      "sudo sh -c 'apt-get install -y asterisk sox'",
      "sudo sh -c 'mkdir -p /usr/src/codecs/opus'",
      "sudo sh -c 'cd /usr/src/codecs/opus && curl -sL http://downloads.digium.com/pub/telephony/codec_opus/${local.opus_codec}.tar.gz | tar --strip-components 1 -xz'",
      "sudo sh -c 'cp /usr/src/codecs/opus/*.so /usr/lib/x86_64-linux-gnu/asterisk/modules/'",
      "sudo sh -c 'cp /usr/src/codecs/opus/codec_opus_config-en_US.xml /usr/share/asterisk/documentation/'",
      "sudo sh -c 'rm /usr/lib/x86_64-linux-gnu/asterisk/modules/format_ogg_opus_open_source.so'",
      "mkdir /tmp/asterisk/",

    ]
  }
  provisioner "file" { 
	source = "conf"
	destination = "/tmp/asterisk/"
  }
  provisioner "file" { 
	source = "scripts"
	destination = "/tmp/asterisk/"
  }
  provisioner "shell" {
    inline = [
      "sudo sh -c 'cp -v /tmp/asterisk/conf/* /etc/asterisk/'",
      "sudo sh -c 'cp -v /tmp/asterisk/scripts/asterisk_network_setup.sh /usr/sbin/'",
      "sudo sh -c 'chmod a+x /tmp/asterisk/scripts/asterisk_network_setup.sh'",
      "sudo sh -c 'mv /tmp/asterisk/scripts/asterisk.service /etc/systemd/system'",
      "sudo sh -c 'chmod 644 /etc/systemd/system/asterisk.service'",
    ]
  }
}

