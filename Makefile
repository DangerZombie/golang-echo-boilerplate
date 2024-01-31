init:
	go mod tidy

clean:
	find . -name "*.mock.gen.go" -type f -delete

# Generate mock files
GO_FILES := $(shell find repository service helper/auth endpoint -name "service.go" -o -name "repository.go" -o -name "auth.go" -o -name "endpoint.go")
GEN_GO_FILES := $(GO_FILES:%.go=%.mock.gen.go)

generate_mocks: $(GEN_GO_FILES)
$(GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

generate_swagger:
	swagger generate spec -m -o ./swagger.yaml

generate_test_coverage:
	go tool cover -html=coverage.out

test_unit:
	go test -short -coverprofile=coverage.out -v ./... -coverpkg=./...

test_integration:
	go test -coverprofile=coverage_integration.out -v ./... -tags=integration