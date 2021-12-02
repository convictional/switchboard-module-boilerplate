#!/bin/bash

cd triggers/gcp/
gcloud --project switchboard-demo functions deploy triggerPubSub \
	--trigger-topic getProduct --runtime go116
cd ../../