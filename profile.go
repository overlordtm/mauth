package mauth

import "golang.org/x/oauth2"

type Profile interface {
	FirstName() string
	LastName() string
	Email() string
	Provider() string
	Token() *oauth2.Token
}
