.PHONY: all

pre: ## fetch the dependencies
	@go get -u github.com/goadesign/goa/...
	@curl https://glide.sh/get | sh

fmt: ## format the source code
	gofmt -w $(GOFMT_FILES)	

build:
	go build -o infra

gen:
	rm -f main.go
	goagen bootstrap -d github.com/perigee/terrant/design	
