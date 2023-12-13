
GOPATH:=$(shell go env GOPATH)
MODIFY=proto/

.PHONY: proto
proto:
    
	protoc --micro_out=${MODIFY} --go_out=${MODIFY} ${MODIFY}order-api.proto
    

.PHONY: build
build: proto

	go build -o order-api-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t order-api-api:latest