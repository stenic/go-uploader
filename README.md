# Go-Uploader

Go-uploader is tool allowing you to upload a file to a set of services.

## Installation

```shell
# homebrew
brew install stenic/tap/go-uploader

# gofish
gofish rig add https://github.com/stenic/fish-food
gofish install github.com/stenic/fish-food/go-uploader

# scoop
scoop bucket add go-uploader https://github.com/stenic/scoop-bucket.git
scoop install go-uploader

# go
go install github.com/stenic/go-uploader@latest

# docker 
docker pull ghcr.io/stenic/go-uploader:latest

# dockerfile
COPY --from=ghcr.io/stenic/go-uploader:latest /go-uploader /usr/local/bin/
```

> For even more options, check the [releases page](https://github.com/stenic/go-uploader/releases).


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
