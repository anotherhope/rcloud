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
	for os in $(TARGET_OS);do                                                                                          \
		for arch in $(TARGET_ARCH);do                                                                                  \
			echo $$os/$$arch;                                                                                          \
			env GOOS=$$os GOARCH=$$arch go build -o $(CURDIR)/.build/rcloud-$$os-$$arch ./app/main.go ;                \
			md5sum $(CURDIR)/.build/rcloud-$$os-$$arch | awk '{print $$1}' > $(CURDIR)/.build/rcloud-$$os-$$arch.md5;  \
			sha1sum $(CURDIR)/.build/rcloud-$$os-$$arch | awk '{print $$1}' > $(CURDIR)/.build/rcloud-$$os-$$arch.sha1; \
		done;                                                                                           		 	   \
	done;

build: vendor ## Build project for local
	go build -o $(CURDIR)/.build/rcloud ./app/main.go

run: ## Run without build project
	go run ./app/main.go

vendor: update-go-deps
	go mod tidy
	go mod vendor

help: #Pour générer automatiquement l'aide ## Display all commands available
	$(eval PADDING=$(shell grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk '{ print length($$1)-1 }' | sort -n | tail -n 1))
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-$(PADDING)s\033[0m %s\n", $$1, $$2}'

update-go-deps:
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done

serve:  ## Serve Markdown for preview
	npm install
	npm start

test:
#rm -rf ../bench
#mkdir -p ../bench
	number=1 ; while [[ $$number -le 100 ]] ; do \
        echo $$number > ../bench/file-$$number.txt; \
        ((number = number + 1)) ; \
    done

kill:
	ps aux | grep "rclone" | grep -v grep | awk '{print $$2}' | xargs kill -9