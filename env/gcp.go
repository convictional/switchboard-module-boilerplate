package env

import "os"

func GCP_PROJECT() string {
	return os.Getenv("GOOGLE_CLOUD_PROJECT")
}

func GCP_LOAD_TOPIC_NAME() string {
	return os.Getenv("GCP_LOAD_TOPIC_NAME")
}