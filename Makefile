.PHONY: help image testing
.SILENT: 
.DEFAULT_GOAL = help

PROJECT_NAME = $(shell basename $(CURDIR))

#TARGET_OS = "windows" "linux" "darwin" "openbsd"
TARGET_OS = "linux" "darwin" "openbsd"
TARGET_ARCH = "amd64" "arm64" 

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)

clean:
	rm -rf $(CURDIR)/.build/*

cross-build: clean vendor  ## Build project for all supported platform
	for os in $(TARGET_OS);do                                                                           \
		for arch in $(TARGET_ARCH);do                                                                   \
			echo $$os/$$arch;                                                                           \
			env GOOS=$$os GOARCH=$$arch go build -o $(CURDIR)/.build/rcloud-$$os-$$arch ./app/main.go ; \
		done;                                                                                           \
	done;

build: vendor ## Build project for local
	go build -o $(CURDIR)/.build/rcloud ./app/main.go

run: ## Run without build project
	go run ./app/main.go

vendor:
	go mod vendor
	go mod tidy

help: #Pour générer automatiquement l'aide ## Display all commands available
	$(eval PADDING=$(shell grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk '{ print length($$1)-1 }' | sort -n | tail -n 1))
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-$(PADDING)s\033[0m %s\n", $$1, $$2}'

#  /opt/homebrew/Library/Taps/homebrew/homebrew-core/