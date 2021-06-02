.PHONY:
.SILENT:
.DEFAULT_GOAL := run

run:
	echo "RUN"

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

lint:
	golangci-lint run