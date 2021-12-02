#!/bin/bash

gcloud --project switchboard-demo functions deploy httpTriggerEvent \
	--trigger-http --security-level=secure-optional --entry-point=HttpTriggerEvent --runtime=go116