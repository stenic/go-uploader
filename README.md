# Go-Uploader
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fstenic%2Fgo-uploader.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fstenic%2Fgo-uploader?ref=badge_shield)


Go-uploader is tool allowing you to upload a file to a set of services.

## Install

```sh
wget https://github.com/stenic/go-uploader/releases/latest/download/go-uploader_$GOOS_$GOARCH.gz
gunzip go-uploader_*.gz
chmod +x go-uploader_*
mv go-uploader_* /usr/local/bin/go-uploader
```

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


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fstenic%2Fgo-uploader.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fstenic%2Fgo-uploader?ref=badge_large)