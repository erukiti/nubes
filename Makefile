build: build-mruby
	@CGO_CFLAGS="-I`pwd`/mruby/include" CGO_LDFLAGS="-L`pwd`/mruby/build/host/lib/ -lmruby -lm" go build

build-mruby:
	@cd mruby && MRUBY_CONFIG=../task/build_config.rb ruby ./minirake

clean:
	@go clean
	@(cd mruby && make clean)