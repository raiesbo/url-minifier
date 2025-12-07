package main

import (
	"github.com/raiesbo/servertools"
	"github.com/raiesbo/url-minifier/internal/models"
	"github.com/raiesbo/url-minifier/internal/url"
)

type application struct {
	url   *url.Service
	users *models.UserModel
	servertools.Tools
}
