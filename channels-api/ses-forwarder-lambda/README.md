# AWS lambda to forward email from SES

This will take an e-mail from amazon SES and invoke the mail-api service
## Installation
1. [validate your domain](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-verification.html) in aws ses
1. create an action chain that does as follows
    * insert e-mail into s3 bucket (create s3 bucket)
    * invoke this lamda
1. add IAM permissions to the lamda role to allow access to the S3 bucket
1. set the envionment variables of the lambda suit your envionment


## Configuring the lambda
the lamda takes 3 envionment variables

| variable          | meaning                                             |
|-------------------|-----------------------------------------------------|
| OL_MAIL_CB_HOST   | ip or hostname of the channel-api server            |
| OL_MAIL_CB_PORT	  | port channel-api server is listening on             |
| OL_MAIL_S3_BUCKET | name of the s3 bucket you created for mail storage  |
| OL_API_KEY        | API key that will be sent to the channel-api server |


## Permissions for the lamda
Add the following to the Statement section of the policy

Change YOUR_S3_BUCKET to be the name of the S3 bucket you created
```json
        {
            "Effect": "Allow",
            "Action": [
                "s3:*"
            ],
            "Resource": [
                "arn:aws:s3:::YOUR_S3_BUCKET",
                "arn:aws:s3:::YOUR_S3_BUCKET/*"
            ]
        }
```