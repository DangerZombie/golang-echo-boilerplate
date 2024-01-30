init:
	go mod tidy

clean:
	find . -name "*.mock.gen.go" -type f -delete

# Generate mock files
GO_FILES := $(shell find repository service helper/auth -name "service.go" -o -name "repository.go" -o -name "auth.go")
GEN_GO_FILES := $(GO_FILES:%.go=%.mock.gen.go)

generate_mocks: $(GEN_GO_FILES)
$(GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))

test_unit:
	go test -short -coverprofile coverage.out -v ./... -coverpkg=./...

test_integration:
	go test -coverprofile coverage_integration.out -v ./... -tags=integration