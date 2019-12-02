BIN_DIR = bin
TOOLS_DIR = $(BIN_DIR)/tools

DAYS = $(shell find -maxdepth 1 -name "day*" | grep -o '[0-9]\+$$')

export PATH := $(shell pwd)/$(TOOLS_DIR):$(PATH)

.PHONY: all
all:
	@echo "run specific day with 'make XX'"

.PHONY: $(DAYS)
$(DAYS):
	./scripts/run.sh $@

$(addsuffix /input,$(addprefix day,$(DAYS))):
	./scripts/input.sh $(shell echo $@ | grep -o '[1-9]\+[0-9]*')

.PHONY: tools
tools:
	mkdir -p $(TOOLS_DIR)
	grep _ < tools.go \
		| cut --delimiter='"' --fields=2 \
		| xargs --verbose env GOBIN=$$(realpath $(TOOLS_DIR)) go install

.PHONY: test
test:
	go test ./...
	golint -set_exit_status ./...

.PHONY: generate
generate:
	go generate ./...
