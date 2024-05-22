package webserver

import (
	"net/http"
	"time"
)

type IServer interface {
	Run()
}

type ServerContext struct {
	Handler http.Handler
	Host    string

	CertFile interface{}
	KeyFile  interface{}

	Timeout      time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}
