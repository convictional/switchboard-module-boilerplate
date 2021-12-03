#!/bin/bash

./scripts/package-lambda.sh

aws lambda update-function-code --function-name DemoSwitchboard \
--zip-file fileb://function.zip