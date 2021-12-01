package env

import "os"

func IsBuyer() bool {
	return os.Getenv("IS_BUYER") == "true"
}

func ConvictionalAPIKey() string {
	return os.Getenv("CONVICTIONAL_API_KEY")
}