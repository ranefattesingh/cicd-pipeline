COVER_FILE=coverage.out
COVER_HTML=coverage.html

test:
	go test ./...

test.coverage:
	ENVIRONMENT=test go test $(shell go list ./... | grep -v "vendor" | grep -v "integration") -p=2 -v
	@go list ./... | grep -v "vendor" | grep -v "integration" | xargs go test -count 1 -cover -short -race -timeout 1m -coverprofile ${COVER_FILE}
	@go tool cover -html=${COVER_FILE} -o ${COVER_HTML}
	@go tool cover -func ${COVER_FILE} | tail -1 | xargs echo test coverage:

lint:
	golangci-lint run -c .golangci.yml

build-deps:
	go mod download

build:
	go build -o app main.go

start:
	./app