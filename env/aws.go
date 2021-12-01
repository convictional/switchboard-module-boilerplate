package env

import "os"

func LoadSQS() string {
	return os.Getenv("LOAD_SQS")
}