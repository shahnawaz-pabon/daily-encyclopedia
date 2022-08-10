<div align="center">
  <a href="https://www.dynamodbguide.com/">
    <img height="96" width="96" alt="dynamodb" src="../logos/dynamodb.svg"/>
  </a>
  <h1>DynamoDB</h1>
</div>

# Table of Contents

- [Create dynamodb client](#create-dynamodb-client)
- [Resources](#resources)

## Create dynamodb client

```python
from starlette.config import Config
import boto3

config = Config(".env")

# Make sure you have .env file having these keys
AWS_ACCESS_KEY_ID: str = config("AWS_ACCESS_KEY_ID", cast=str)
AWS_SECRET_ACCESS_KEY: str = config("AWS_SECRET_ACCESS_KEY", cast=str)
REGION: str = config("REGION", cast=str)

dynamodb_client = boto3.client(
        "dynamodb",
        aws_access_key_id=AWS_ACCESS_KEY_ID,
        aws_secret_access_key=AWS_SECRET_ACCESS_KEY,
        region_name=REGION,
    )
```

<br>

## Resources

- [Boto3 dynamodb official docs][boto3-dynamodb]

<!-- Definitions -->

[boto3-dynamodb]: https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/dynamodb.html#client
