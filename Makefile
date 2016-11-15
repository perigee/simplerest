.PHONY: all


fmt: ## format the source code
	gofmt -w $(GOFMT_FILES)	

build:
	docker build --build-arg proxy=$(http_proxy) -t simpleservice/testing .
