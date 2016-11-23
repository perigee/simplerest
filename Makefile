.PHONY: all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
.DEFAULT_GOAL := build

get:
	go get -d -u ./...

pre: ## fetch the dependencies
	@go get -u github.com/goadesign/goa/...
	@curl https://glide.sh/get | sh

fmt: ## format the source code
	gofmt -w $(GOFMT_FILES)	

build:
	go build -o infra

run: build
	./infra


gen:
	goagen app -d github.com/perigee/terrant/design
	goagen swagger -d github.com/perigee/terrant/design
	goagen schema -d github.com/perigee/terrant/design
	goagen client -d github.com/perigee/terrant/design
	goagen js -d github.com/perigee/terrant/design


bootstrap:
	rm -f main.go
	goagen bootstrap -d github.com/perigee/terrant/design

docker:
	docker build --build-arg proxy=$(http_proxy) -t perigee/terrant:dev -f Dockerfile_dev .
