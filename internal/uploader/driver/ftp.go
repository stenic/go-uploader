package driver

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"time"

	ftp "github.com/jlaffaye/ftp"
	logger "github.com/stenic/go-uploader/internal/uploader/log"
)

func init() {
	AddDriver("ftp", &FtpDriver{})
}

type FtpDriver struct {
}

func (d *FtpDriver) Upload(ctx context.Context, src *url.URL, dst *url.URL) (*UploadResult, error) {
	log := logger.LoggerFromContext(ctx)

	var srcPath = src.Host + src.Path
	var dstPath = dst.Host + dst.Path

	ftpClient, err := ftp.Dial(dst.Host, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}
	defer ftpClient.Quit()
	log.Debug("Connection completed")

	pass, _ := dst.User.Password()
	if err = ftpClient.Login(dst.User.Username(), pass); err != nil {
		return nil, err
	}
	log.Debug("Login completed")

	if err = ftpClient.ChangeDir(filepath.Dir(dst.Path)); err != nil {
		panic(err)
	}
	log.Debug("Change Dir completed")

	file, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	log.Debug("Open file completed")

	if err = ftpClient.Stor(filepath.Base(dstPath), file); err != nil {
		panic(err)
	}
	log.Debug("Upload completed")

	return nil, err
}
