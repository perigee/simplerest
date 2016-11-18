.PHONY: all

pre: ## fetch the dependencies
	@go get -u github.com/goadesign/goa/...
	@curl https://glide.sh/get | sh

fmt: ## format the source code
	gofmt -w $(GOFMT_FILES)	

build:
	docker build --build-arg proxy=$(http_proxy) -t simpleservice/testing .

gen:
	rm -f main.go
	goagen bootstrap -d github.com/perigee/terrant/design	