# Go-Uploader

Go-uploader is tool allowing you to upload a file to a set of services.

## Drivers

### Local

```sh
go-uploader upload examples/assets/testfile.txt local://.tmp/local/testfile.txt
```

### Ftp

```sh
go-uploader upload examples/assets/testfile.txt ftp://test:test@localhost:21/testfile.txt
```

### Minio

```sh
go-uploader upload examples/assets/testfile.txt minio://test:testtest@127.0.0.1:9000/bucket/testfile.txt
```

### Http/Https

```sh
go-uploader upload examples/assets/testfile.txt http+post://127.0.0.1:4444/files/testfile.txt
```
