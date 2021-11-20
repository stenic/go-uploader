package driver

import (
	"context"
	"net/url"

	cp "github.com/nmrshll/go-cp"
	logger "github.com/stenic/go-uploader/internal/uploader/log"
)

func init() {
	AddDriver("local", &LocalDriver{})
}

type LocalDriver struct {
}

func (d *LocalDriver) Upload(ctx context.Context, src *url.URL, dst *url.URL) (*UploadResult, error) {
	log := logger.LoggerFromContext(ctx)

	var srcPath = src.Host + src.Path
	var dstPath = dst.Host + dst.Path
	log.Debugf("Copy %s to %s\n", srcPath, dstPath)

	err := cp.CopyFile(srcPath, dstPath)

	return nil, err
}
