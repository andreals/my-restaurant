.PHONY: run
run:
	@go run cmd/cli/main.go 

.PHONY: build
build:
	@go build -o ./bin/cli ./cmd/cli

.PHONY: test
test:
	@go test -race ./...

.PHONY: docker-build
docker-build:
	@docker build -t restaurant . -f ./build/Dockerfile

.PHONY: docker-run
docker-run:
	@docker run -it restaurant
