package driver

import (
	"context"
	"fmt"
	"net/url"
)

type UploadResult struct{}

type Driver interface {
	Upload(context.Context, *url.URL, *url.URL) (*UploadResult, error)
}

var driverRegistry = map[string]Driver{}

func GetDriver(name string) (Driver, error) {
	if val, ok := driverRegistry[name]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("Driver %s could not be found", name)
}

func AddDriver(name string, driver Driver) {
	driverRegistry[name] = driver
}
