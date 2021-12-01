# GCP deploy makefile

.PHONY: deploy-gcp-pubsub

deploy-gcp-pubsub:
	gcloud --project switchboard-demo functions deploy triggerPubSub \
	--trigger-topic getProduct --runtime go116

.PHONY: create-gcp-pubsub

create-gcp-pubsub:
	gcloud --project switchboard-demo pubsub topics create getProduct

