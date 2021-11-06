test:
	rm -rf .tmp
	mkdir -p .tmp/{ftp,local}
	docker-compose up -d
	sleep 10
	go run main.go upload local://examples/assets/testfile.txt local://.tmp/local/testfile.txt
	go run main.go upload local://examples/assets/testfile.txt local:///$(PWD)/.tmp/local/testfile_abs.txt
	go run main.go upload local://examples/assets/testfile.txt ftp://test:test@localhost:21/testfile.txt
	go run main.go upload local://examples/assets/testfile.txt minio://test:testtest@127.0.0.1:9000/somebucketname/testfile2.txt
