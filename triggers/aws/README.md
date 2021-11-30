# AWS

## Zipping for Lambda

```
cd triggers/aws/
GOOS=linux GOARCH=amd64 go build -o main .
zip function.zip main
# Upload this file to Lambda
# Make sure your handler is set to main within Lambda settings
```