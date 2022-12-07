DAYS = $(shell find -maxdepth 1 -name "day*" | grep -o '[0-9]\+$$')

export AOC_TOKEN = $(shell cat session-token)

.PHONY: all
all:
	go test ./...

.PHONY: $(DAYS)
$(DAYS):
	go test ./day$@

day%: day%/input

day%/input:
	mkdir -p $(shell dirname $@)
	go run ./dl -token "$$AOC_TOKEN" -day "$*" 
