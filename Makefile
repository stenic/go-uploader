INFO=\n\033[0;33m
NC=\033[0m\n

test: services
	@printf "${INFO}Build binary${NC}"
	go build -ldflags="`vembed`" -o .tmp/go-uploader main.go
	ls -lah .tmp

	@printf "${INFO}Testing singles${NC}"
	.tmp/go-uploader upload examples/assets/testfile.txt local://.tmp/local/testfile.txt
	.tmp/go-uploader upload examples/assets/testfile.txt local:///$(PWD)/.tmp/local/testfile_abs.txt
	.tmp/go-uploader upload examples/assets/testfile.txt ftp://test:test@localhost:21/testfile.txt
	.tmp/go-uploader upload examples/assets/testfile.txt minio://test:testtest@127.0.0.1:9000/somebucketname/testfile.txt
	.tmp/go-uploader upload examples/assets/testfile.txt http+post://127.0.0.1:4444/files/testfile.txt

	@printf "${INFO}Testing multiple${NC}"
	.tmp/go-uploader upload examples/assets/testfile.txt local://.tmp/local/testfile.txt \
		local:///$(PWD)/.tmp/local/testfile_abs.txt \
		ftp://test:test@localhost:21/testfile.txt \
		minio://test:testtest@127.0.0.1:9000/somebucketname/testfile2.txt \
		http+post://127.0.0.1:4444/files/testfile.txt

services:
	@printf "${INFO}Preparing setup${NC}"
	rm -rf .tmp
	mkdir -p .tmp/{ftp,local}
	docker-compose up -d

clean:
	rm -rf dist .tmp

buildPack = GOOS=$(1) GOARCH=$(2) go build -ldflags="`vembed`" -o dist/go-uploader$(3) main.go \
	&& (cd dist; zip $(1)_$(2).zip go-uploader$(3); rm go-uploader$(3))

dist:
	$(call buildPack,windows,amd64,.exe)
	$(call buildPack,windows,arm64,.exe)
	$(call buildPack,darwin,amd64,)
	$(call buildPack,darwin,arm64,)
	$(call buildPack,linux,amd64,)
	$(call buildPack,linux,arm64,)
