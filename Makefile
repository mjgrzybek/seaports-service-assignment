test: mocks
	go test ./...

mocks:
	go install github.com/vektra/mockery/v2@latest
	mockery --all