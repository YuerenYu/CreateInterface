package gracehttp

/**
	this package is quote https://github.com/tabalt/gracehttp
	I changed a method, so did't use "import"
 */

import (
	"net/http"
	"time"
)

const (
	DEFAULT_READ_TIMEOUT  = 60 * time.Second
	DEFAULT_WRITE_TIMEOUT = DEFAULT_READ_TIMEOUT
)

// refer http.ListenAndServe
func ListenAndServe(addr string, handler http.Handler) error {
	return NewServer(addr, handler, DEFAULT_READ_TIMEOUT, DEFAULT_WRITE_TIMEOUT).ListenAndServe()
}

// refer http.ListenAndServeTLS
func ListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error {
	return NewServer(addr, handler, DEFAULT_READ_TIMEOUT, DEFAULT_WRITE_TIMEOUT).ListenAndServeTLS(certFile, keyFile)
}
