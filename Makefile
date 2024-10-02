CLOUD_API_VERSION := $(or ${CLOUD_API_VERSION},$(shell curl -L "https://raw.githubusercontent.com/fi-ts/releases/develop/release.yaml" | yq e '.docker-images.metal-stack.extensions.cloud-api.tag' -))

release:: generate-client mocks gofmt test;

.PHONY: generate-client
generate-client:
	rm -rf api
	mkdir -p api
	yq e -ij ".info.version=\"${CLOUD_API_VERSION}\"" cloud-api.json
	yq e '.info.version' cloud-api.json
	docker pull metalstack/builder
	docker run --rm \
		--user $$(id -u):$$(id -g) \
		-v ${PWD}:/work \
		metalstack/builder swagger generate client -f cloud-api.json -t api

.PHONY: mocks
mocks:
	rm -rf test/mocks
	docker run --user $$(id -u):$$(id -g) --rm -w /work -v ${PWD}:/work vektra/mockery:v2.46.1 -r --keeptree --inpackage --dir api/client --output test/mocks --all
	go run ./test/client/generate/generate_mock_client.go

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: test
test:
	go test ./... -coverprofile=coverage.out -covermode=atomic && go tool cover -func=coverage.out

.PHONY: golangcicheck
golangcicheck:
	@/bin/bash -c "type -P golangci-lint;" 2>/dev/null || (echo "golangci-lint is required but not available in current PATH. Install: https://github.com/golangci/golangci-lint#install"; exit 1)

.PHONY: lint
lint: golangcicheck
	CGO_ENABLED=1 golangci-lint run
