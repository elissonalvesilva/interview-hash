
.PHONY: build start stop logs-checkout-service logs-analyzer logs-tail tests-docker tests-local

CONTAINER_NAME_API := checkout-service

build:
	docker-compose build

start:
	docker-compose up -d

stop:
	docker-compose down

logs-checkout-service:
	docker logs -f $(CONTAINER_NAME_API)

logs-tail:
	docker logs -f --tail 100 $(CONTAINER_NAME_API)

tests:
	go test ./tests/... -coverpkg=./my-ecommerce/... -cover -coverprofile=coverage.out

show-coverage-html:
	go tool cover -html=coverage.out