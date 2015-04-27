// +build appengine
package mauth

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func GetContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}
