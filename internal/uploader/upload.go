package uploader

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stenic/go-uploader/internal/uploader/driver"
	logger "github.com/stenic/go-uploader/internal/uploader/log"
)

var cmdUpload = &cobra.Command{
	Use:   "upload SRC DEST",
	Short: "Do the upload",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := logger.ContextWithLogger(context.Background(), logrus.WithFields(logrus.Fields{}))
		src, err := url.Parse(args[0])
		if err != nil {
			return err
		}

		var wg sync.WaitGroup
		for _, d := range args[1:] {
			wg.Add(1)
			dArg := d
			go func() {
				defer wg.Done()
				doUpload(ctx, src, dArg)
			}()
		}
		wg.Wait()

		return nil
	},
}

func doUpload(ctx context.Context, src *url.URL, dstArg string) error {
	log := logger.LoggerFromContext(ctx)

	dst, err := url.Parse(dstArg)
	if err != nil {
		return err
	}

	if dst.Scheme == "" {
		return errors.New("DST needs to contain a schema (local://path)")
	}

	log.Debugf("Uploading %s to %v (protocol %s) \n", src.Host+src.Path, dst.Host+dst.Path, dst.Scheme)

	driver, err := driver.GetDriver(dst.Scheme)
	if err != nil {
		return err
	}

	log.Debugf("Found driver for %s %T\n", dst.Scheme, driver)

	start := time.Now()
	_, err = driver.Upload(
		logger.ContextWithLogger(ctx, logrus.WithFields(logrus.Fields{"driver": dst.Scheme})),
		src,
		dst,
	)
	duration := time.Since(start)

	log.Infof("%s: Upload completed in %f seconds", dst.Scheme, duration.Seconds())

	return err

}

func init() {
	rootCmd.AddCommand(cmdUpload)
}
