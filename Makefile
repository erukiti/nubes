build-mruby:
	cd mruby && MRUBY_CONFIG=../task/build_config.rb ruby ./minirake

build: build-mruby
	go build

fmt:
	go fmt ./...
