build:
	docker build -f build/Dockerfile . -t seaport-service

test: mocks
	go test ./...

mocks:
	go install github.com/vektra/mockery/v2@latest
	mockery --all

.PHONY: build test mocks