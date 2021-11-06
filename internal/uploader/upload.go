package uploader

import (
	"context"
	"errors"
	"net/url"

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
		ctx := context.Background()
		log := logrus.New()

		src, err := url.Parse(args[0])
		if err != nil {
			return err
		}
		dst, err := url.Parse(args[1])
		if err != nil {
			return err
		}
		if dst.Scheme == "" {
			return errors.New("DST needs to contain a schema (local://path)")
		}

		log.Infof("Going to upload %s to %v protocol %s \n", args[0], dst.Host+dst.Path, dst.Scheme)

		driver, err := driver.GetDriver(dst.Scheme)
		if err != nil {
			return err
		}

		res, err := driver.Upload(
			logger.ContextWithLogger(ctx, logrus.WithFields(logrus.Fields{"driver": "local"})),
			src,
			dst,
		)

		log.Debugf("%v\n", res)

		return err
	},
}

func init() {
	rootCmd.AddCommand(cmdUpload)
}
