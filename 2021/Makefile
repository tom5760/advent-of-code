DAYS = $(shell find -maxdepth 1 -name "day*" | grep -o '[0-9]\+$$')

export AOC_TOKEN = $(shell cat session-token)

.PHONY: all
all:
	@echo "run specific day with 'make XX'"

.PHONY: $(DAYS)
$(DAYS):
	gotip test ./day$@
	gotip run ./day$@ < day$@/input

day%: day%/input

day%/input:
	mkdir -p $(shell dirname $@)
	gotip run ./dl -day $*
