.PHONY: all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)


get:
	go get -d -u ./...

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

docker:
	docker build --build-arg proxy=$(http_proxy) -t perigee/terrant .
