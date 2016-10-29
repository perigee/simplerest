build:
	docker build --build-arg proxy=$(http_proxy) -t simpleservice/testing .
proto:
	protoc -I pb/ pb/service.proto --go_out=plugins=grpc:pb
