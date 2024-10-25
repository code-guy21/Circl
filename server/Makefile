#Makefile

# Variables

BINARY_NAME = pingup-api
MAIN_FILE = cmd/api/main.go

# Developement
.PHONY: run
run:
	go run ${MAIN_FILE}

.PHONY: build
build:
	go build -o bin/${BINARY_NAME} ${MAIN_FILE}
	
.PHONY: clean
clean:
	go clean
	rm -f bin/${BINARY_NAME}

.PHONY: test
test:
	go test ./...

.PHONY: tidy
tidy:
	go mod tidy

# Docker
.PHONY: docker-build
docker-build:
	docker build -t ${BINARY_NAME}

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 ${BINARY_NAME}