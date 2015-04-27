package mauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type fbProfile struct {
	FBFirstName string        `json:"first_name"`
	FBLastName  string        `json:"last_name"`
	FBEmail     string        `json:"email"`
	FBToken     *oauth2.Token `json:"-"`
}

func (p *fbProfile) FirstName() string {
	return p.FBFirstName
}
func (p *fbProfile) LastName() string {
	return p.FBLastName
}
func (p *fbProfile) Email() string {
	return p.FBEmail
}
func (p *fbProfile) Token() *oauth2.Token {
	return p.FBToken
}
func (p *fbProfile) Provider() string {
	return "facebook"
}

func appsecretProof(clientSecret, accessToken string) string {
	hash := hmac.New(sha256.New, []byte(clientSecret))
	hash.Write([]byte(accessToken))
	return hex.EncodeToString(hash.Sum(nil))
}

func getFacebookProfile(conf *Config, c *http.Client, t *oauth2.Token) *fbProfile {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://graph.facebook.com/me?appsecret_proof=%s", appsecretProof(conf.OAuth2.ClientSecret, t.AccessToken)), nil)
	if err != nil {
		panic(err)
	}

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	profile := &fbProfile{}
	if err := json.NewDecoder(res.Body).Decode(profile); err != nil {
		panic(err)
	}

	profile.FBToken = t

	return profile
}
