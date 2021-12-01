# GCP deploy makefile

.PHONY: deploy-gcp-pubsub

deploy-gcp-pubsub:
	gcloud --project switchboard-demo functions deploy triggerPubSub \
	--trigger-topic getProduct --runtime go116

deploy-gcp-http:
	gcloud --project switchboard-demo functions deploy httpTriggerEvent \
	--trigger-http --security-level=secure-optional

.PHONY: create-gcp-pubsub

create-gcp-pubsub:
	gcloud --project switchboard-demo pubsub topics create getProduct

