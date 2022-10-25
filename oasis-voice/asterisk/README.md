# openline-voice-asterisk-server

Currently there is only a packer ami image available

to build the api image you must have AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY AWS_REGION properly set with your AWS credentials

then you can run the following commands to build the image

```
packer init aws-ubuntu.pkr.hcl
packer validate aws-ubuntu.pkr.hcl
packer build aws-ubuntu.pkr.hcl
```
