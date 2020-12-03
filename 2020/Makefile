DAYS = $(shell find src/bin -maxdepth 1 -name "day*" | sed 's/[^0-9]//g')

.PHONY: all
all:
	@echo "run specific day with 'make XX'"

.PHONY: $(DAYS)
$(DAYS):
	./scripts/run.sh $@

$(addprefix inputs/day,$(DAYS)):
	./scripts/input.sh $(shell echo $@ | grep -o '[1-9]\+[0-9]*')

.PHONY: test
test:
	cargo test