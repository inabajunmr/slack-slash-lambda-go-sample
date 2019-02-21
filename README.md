# Slack Slash Commands sample

Slack Slash Commands sample implementation by golang and AWS Lambda.

## Requirements

* AWS CLI already configured with Administrator permission
* [aws-sam-cli](https://github.com/awslabs/aws-sam-cli)
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Deploy

### Config

You must set environment variables VERIFICATION_TOKEN on template.yaml.

#### Example
```yaml
Variables:
  VERIFICATION_TOKEN: xxxxx
```

### Deploy Command
```
make deploy BUCKET="your S3 bucket name" PROFILE="your aws profile"
```
