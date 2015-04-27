// +build !appengine

package mauth

import (
	"net/http"

	"golang.org/x/net/context"
)

func GetContext(r *http.Request) context.Context {
	return context.TODO()
}
