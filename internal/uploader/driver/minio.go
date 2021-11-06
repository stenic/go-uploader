package driver

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	logger "github.com/stenic/go-uploader/internal/uploader/log"
)

func init() {
	AddDriver("minio", &MinioDriver{})
}

type MinioDriver struct {
}

func (d *MinioDriver) Upload(ctx context.Context, src *url.URL, dst *url.URL) (*UploadResult, error) {
	log := logger.LoggerFromContext(ctx)

	var srcPath = src.Host + src.Path

	pass, _ := dst.User.Password()
	client, err := minio.New(dst.Host, &minio.Options{
		Creds: credentials.NewStaticV4(dst.User.Username(), pass, ""),
	})
	if err != nil {
		return nil, err
	}
	log.Debug("Client created")

	file, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	log.Debug("Open file completed")

	bucket, path := d.getBucketAndPath(dst.Path)
	log.Debugf("Found bucket '%s' and path '%s'", bucket, path)

	_, err = client.PutObject(ctx, bucket, path, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

	return nil, err
}

func (d *MinioDriver) getBucketAndPath(path string) (string, string) {
	bucket := path
	for i := strings.Count(path, "/"); i > 1; i-- {
		bucket = filepath.Dir(bucket)
	}
	return bucket[1:], path[len(bucket)+1:]
}
