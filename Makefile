.PHONY: all
all: vendor dockerbuild containerup

vendor:
	@go mod download
	@go mod vendor

dockerbuild:
	@echo "...building image iata-finder"
	@docker build -t abelgoodwin1988/iata-finder:latest . --no-cache

containerup:
	@echo "...composing container iata-finder"
	@sh -c  "docker-compose up"

.PHONY: rpc
rpc:
	@echo "...generating .pb.go file"
	@protoc rpc/iatafinder.proto --go_out=plugins=grpc:.