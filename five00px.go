// Package five00px provides main 500px API implementation
package five00px

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/mrjones/oauth"
	"github.com/toqueteos/webbrowser"
)

var mainAPIUrl string = "https://api.500px.com/v1/oauth/"

type oAuthInfo struct {
	consumerKey    string
	consumerSecret string
	oauthToken     string
	oauthVerifier  string
}

type Five00px struct {
	oAuthInfo
	c    *oauth.Consumer
	Port uint
}

func genOAuthConsumer(consumerKey, consumerSecret string) *oauth.Consumer {
	return oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   mainAPIUrl + "request_token",
			AuthorizeTokenUrl: mainAPIUrl + "authorize",
			AccessTokenUrl:    mainAPIUrl + "access_token",
		})
}

func New(consumerKey, consumerSecret string) *Five00px {
	return &Five00px{
		oAuthInfo: oAuthInfo{
			consumerKey:    consumerKey,
			consumerSecret: consumerSecret},
		c:    genOAuthConsumer(consumerKey, consumerSecret),
		Port: 8080}
}

func (f00 *Five00px) Debug(enabled bool) {
	f00.c.Debug(enabled)
}

func parseAccessToken(urlQuery string) (oauthToken string, oauthVerifier string, err error) {
	if strings.HasPrefix(urlQuery, "/?") {
		urlQuery = urlQuery[2:]
	}
	val, err := url.ParseQuery(urlQuery)
	if err != nil {
		return "", "", err
	}
	oauthToken = val["oauth_token"][0]
	oauthVerifier = val["oauth_verifier"][0]

	return oauthToken, oauthVerifier, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	oauthToken, oauthVerifier, err := parseAccessToken(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "Token: ", oauthToken, ", Verifier: ", oauthVerifier)
}

func (f00 *Five00px) Auth() {
	_, u, err := f00.c.GetRequestTokenAndUrl(fmt.Sprint("http://127.0.0.1:", f00.Port))
	if err != nil {
		log.Fatal(err)
	}

	webbrowser.Open(u)

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprint(":", f00.Port), nil)
}
