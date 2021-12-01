package main

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_HandleRequest(t *testing.T) {
	t.Run("Demo SQS", func(t *testing.T) {
		sqsRawInput, err := ioutil.ReadFile("./testdata/sqs_input.json")
		assert.Nil(t, err)

		var inputEvent AWSTriggerEvent
		err = json.Unmarshal(sqsRawInput, &inputEvent)
		assert.Nil(t, err)

		os.Setenv("TRANSFORM", "true")
		os.Setenv("LOAD", "true")


		HandleRequest(context.Background(), inputEvent)
	})
}
