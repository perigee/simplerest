build:
	#docker build --build-arg proxy=$(http_proxy) -t simpleservice/testing .
	docker build --build-arg proxy=$(http_proxy) -t simpleservice/terra -f Dockerfile_terraform .
proto:
	protoc -I pb/ pb/service.proto --go_out=plugins=grpc:pb
