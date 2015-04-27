package mauth

import (
	"net/http"

	"github.com/go-martini/martini"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	KeyNextPage  = "next"
	PathError    = "/auth/error"
	CodeRedirect = 302
)

type ContextFn func(r *http.Request) context.Context

type Config struct {
	OAuth2       *oauth2.Config
	Context      ContextFn
	PathError    string
	CodeRedirect int
}

type AuthError struct {
	Error error
}

func Auth(conf *Config) martini.Handler {
	autofillConfig(conf)
	return func(w http.ResponseWriter, r *http.Request) {
		next := extractPath(r.URL.Query().Get(KeyNextPage))
		http.Redirect(w, r, conf.OAuth2.AuthCodeURL(next), conf.CodeRedirect)
	}
}

func Callback(conf *Config) martini.Handler {
	autofillConfig(conf)
	return func(c martini.Context, w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		t, err := conf.OAuth2.Exchange(conf.Context(r), code)

		if err != nil {
			http.Redirect(w, r, conf.PathError, conf.CodeRedirect)
		}

		client := conf.OAuth2.Client(conf.Context(r), t)
		profile := getFacebookProfile(conf, client, t)

		c.MapTo(profile, (*Profile)(nil))
	}
}
