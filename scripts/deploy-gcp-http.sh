#!/bin/bash

cd triggers/gcp/
gcloud --project switchboard-demo functions deploy httpTriggerEvent \
	--trigger-http --security-level=secure-optional
cd ../../