# Help legend
# ##@ = Category
# ## = Target description
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build CLI binary.
	@echo "==> building golang binary"
	version=$$(cat VERSION) && \
	commit=$$(git rev-parse --short HEAD) && \
	now=$$(date +'%Y-%m-%d_%T') && \
	go build \
	-ldflags=" \
	-X 'github.com/droctothorpe/toco/cmd.Version=$$version' \
	-X 'github.com/droctothorpe/toco/cmd.Commit=$$commit' \
	-X 'github.com/droctothorpe/toco/cmd.BuildTime=$$now' " \
	-o toco

.PHONY: install
install: ## Install CLI binary to GOPATH.
	@echo "==> building and installing golang binary"
	version=$$(cat VERSION) && \
	commit=$$(git rev-parse --short HEAD) && \
	now=$$(date +'%Y-%m-%d_%T') && \
	go install \
	-ldflags=" \
	-X 'github.com/droctothorpe/toco/cmd.Version=$$version' \
	-X 'github.com/droctothorpe/toco/cmd.Commit=$$commit' \
	-X 'github.com/droctothorpe/toco/cmd.BuildTime=$$now' " 
	