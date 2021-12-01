cd triggers/aws/
GOOS=linux GOARCH=amd64 go build -o main .
zip function.zip main
cd ../../