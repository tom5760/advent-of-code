DAYS = $(shell find -maxdepth 1 -name "day*" | grep -o '[0-9]\+$$')

.PHONY: all
all:
	@echo "run specific day with 'make XX'"

.PHONY: $(DAYS)
$(DAYS):
	./scripts/run.sh $@

$(addsuffix /input,$(addprefix day,$(DAYS))):
	./scripts/input.sh $(shell echo $@ | grep -o '[1-9]\+[0-9]*')
