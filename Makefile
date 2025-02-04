.PHONY: build push release tag

build:
	docker build -t ghcr.io/haleyrc/honig .

push:
	docker push ghcr.io/haleyrc/honig:latest

release: build tag push
