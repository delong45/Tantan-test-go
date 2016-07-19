package service

import (
	"net/http"
	"tantan/config"
)

type Server struct {
	Conf     *config.Config
	Closing  chan struct{}
	HTTPDone chan struct{}
	Handler  http.Handler
}
