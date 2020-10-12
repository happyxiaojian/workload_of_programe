os := $(if $(GOOS), $(GOOS), $(shell go env GOOS))

txt = hello world
test:
	#这是测试
	@echo $(txt)
	@echo $(os)
a.txt: b.txt c.txt
	cat b.txt c.txt > a.txt

LIST = one two three
all:
	for i in $(LIST); do \
		ehco $$i; \
	done
