CLOUD_API_VERSION := $(or ${CLOUD_API_VERSION},$(shell cat VERSION))

release:: generate-client mocks gofmt test;

.PHONY: generate-client
generate-client:
ifeq ($(CI),true)
	curl -v -LO -H 'Accept: application/vnd.github.v3.raw' -H 'authorization: Bearer $(GENERATE_TOKEN)' https://api.github.com/repos/fi-ts/cloud-api/spec/cloud-api.json
	cat cloud-api.json
else
	rm -rf tmp
	git clone https://github.com/fi-ts/cloud-api.git tmp
	cd tmp && git checkout $(CLOUD_API_VERSION) && cd -
	mv tmp/spec/cloud-api.json .
	rm -rf tmp
endif
	$(MAKE) generate-client-local

.PHONY: generate-client-local
generate-client-local:
	yq e -i -o=json ".info.version=\"${CLOUD_API_VERSION}\"" cloud-api.json
	yq e -o=json '.info.version' cloud-api.json
	rm -rf api
	mkdir -p api
	docker pull ghcr.io/metal-stack/builder
	docker run --rm \
		--user $$(id -u):$$(id -g) \
		-v ${PWD}:/work \
		ghcr.io/metal-stack/builder swagger generate client -f cloud-api.json -t api --struct-tags json

.PHONY: mocks
mocks:
	rm -rf test/mocks
	docker run --rm \
		--user $$(id -u):$$(id -g) \
		-w /work \
		-v ${PWD}:/work \
		vektra/mockery:v2.50.0 -r --keeptree --inpackage --dir api/client --output test/mocks --all
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
