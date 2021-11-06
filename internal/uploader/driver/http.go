package driver

import (
	"context"
	"net/http"
	"net/url"
	"os"
)

func init() {
	AddDriver("http+post", &HttpDriver{
		scheme: "http",
		method: "POST",
	})
	AddDriver("https+post", &HttpDriver{
		scheme: "https",
		method: "POST",
	})
}

type HttpDriver struct {
	scheme string
	method string
}

func (d *HttpDriver) Upload(ctx context.Context, src *url.URL, dst *url.URL) (*UploadResult, error) {
	// log := logger.LoggerFromContext(ctx)
	dst.Scheme = d.scheme

	client := &http.Client{}
	data, err := os.Open(src.Host + src.Path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(d.method, dst.String(), data)
	if err != nil {
		return nil, err
	}
	if dst.User.Username() != "" {
		pass, _ := dst.User.Password()
		req.SetBasicAuth(dst.User.Username(), pass)
	}
	_, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	return nil, err
}
