if [ $1 = "aws" ]
then
  bash ./scripts/package-lambda.sh
else
  echo "Unsupported platform"
  exit 1
fi