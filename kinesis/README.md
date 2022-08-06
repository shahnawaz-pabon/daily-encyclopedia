<div align="center">
  <a href="https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/kinesis.html">
    <img alt="kinesis" src="../logos/aws-kinesis.svg" height="100" width="100"/>
  </a>
  <h1>Kinesis</h1>
</div>

# Table of Contents

- [Create kinesis stream from aws-cli](#create-kinesis-stream-from-aws-cli)
- [Check shard list of a specific stream from aws-cli](#check-shard-list-of-a-specific-stream-from-aws-cli)

<br>

## Create kinesis stream from aws-cli

```bash
$ brew install awscli # installing aws-cli on macOS
$ aws kinesis create-stream \
--stream-name StreamName \
--shard-count 2 \
--region ap-southeast-1
```

<br>

## Check shard list of a specific stream from aws-cli

```bash
$ aws kinesis list-shards --stream-name StreamName
```
