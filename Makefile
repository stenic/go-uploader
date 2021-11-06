test:
	rm -rf .tmp
	mkdir .tmp
	docker-compose up -d
	go run cmd/uploader/main.go upload local://tests/assets/testfile.txt local://.tmp/local/sdf
	
