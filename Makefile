.PHONY: build start stop logs-checkout-service logs-checkout-service logs-tail tests tests-docker

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

tests-docker:
	docker build --no-cache -t $(CONTAINER_NAME_API)-test -f ./Dockerfile.test . && docker run -v ${PWD}:/go/testdir $(CONTAINER_NAME_API)-test

show-coverage-html:
	go tool cover -html=coverage.out