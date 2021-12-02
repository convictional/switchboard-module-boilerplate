# GCP deploy makefile

.PHONY: deploy-gcp-pubsub

deploy-gcp-pubsub:
	./scripts/deploy-gcp-pubsub.sh

deploy-gcp-http:
	./scripts/deploy-gcp-http.sh

.PHONY: create-gcp-pubsub

create-gcp-pubsub:
	gcloud --project switchboard-demo pubsub topics create getProduct

