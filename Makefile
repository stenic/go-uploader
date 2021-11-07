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

dist:
	go get github.com/NoUseFreak/go-vembed/vembed
	go get github.com/mitchellh/gox
	gox -ldflags="`vembed`" -os="linux darwin windows" -arch="amd64 arm64" -output="dist/go-uploader_{{.OS}}_{{.Arch}}"
	(cd dist; gzip *)