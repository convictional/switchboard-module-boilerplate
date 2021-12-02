#!/bin/bash

gcloud --project switchboard-demo functions deploy triggerPubSub \
	--trigger-topic getProduct --runtime go116 --entry-point=TriggerPubSub
