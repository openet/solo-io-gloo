.PHONY: docker-local server
all: server docker-local

# make server is no longer a necessary pre-requisite. The binary is being built in the Dockerfile.
server:
	@CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -a --ldflags '-extldflags "-static"' \
		-tags "netgo" -installsuffix netgo \
		-o server

docker-local: server
	docker build -t gcr.io/solo-public/passthrough-grpc-service-example .

docker-push:
	docker buildx build --no-cache --push --platform linux/arm64,linux/amd64 -t gcr.io/solo-public/passthrough-grpc-service-example .

clean:
	@rm -fr server
