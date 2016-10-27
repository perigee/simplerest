build:
	docker build --build-arg proxy=$(http_proxy) -t simpleservice/testing .
