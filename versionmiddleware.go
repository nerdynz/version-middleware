package versionmiddleware

import (
	"net/http"
	"time"
)

func New() *XVersion {
	return &XVersion{
		version: time.Now().Format(time.Stamp),
	}
}
func Custom(version string) *XVersion {
	return &XVersion{
		version: version,
	}
}

func (v *XVersion) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Set("X-Version", v.version)
	next(rw, r)
}

type XVersion struct {
	version string
}
