# convic-module-boilerplate

## Environment Variables

### General

| Name  | Required | Description   |
| ----- | -------- | ------------- |
| `Example` | True | This is a demo |

### GCP

| Name  | Required | Description   |
| ----- | -------- | ------------- |
| `Example` | True | This is a demo |


### AWS

#### Load

| Name  | Required | Description   |
| ----- | -------- | ------------- |
| `LOAD_SQS` | False | The ARN URL of the SQS you are looking to publish too. You need to make sure your IAM role on the Lambda has publish permissions. Your queue should also allow your IAM role to publish to it. Ex. `https://sqs.us-east-2.amazonaws.com/1234/DemoName` (`https://sqs.<region>.amazonaws.com/<Account number>/<SQS Name>`) |
| `GOOGLE_CLOUD_PROJECT` | False | The GCP project name that this function will run in |
| `GCP_LOAD_TOPIC_NAME` | False | The pub/sub topic name to push messages to from the LOAD switch |

## Running

### Local

#### AWS

```
go run ./triggers/aws/...
```