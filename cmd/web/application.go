package main

import (
	"github.com/raiesbo/servertools"
	"github.com/raiesbo/url-minifier/internal/url"
)

type application struct {
	url *url.Service
	servertools.Tools
}
