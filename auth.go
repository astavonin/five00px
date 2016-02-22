// Package five00px provides ...
package five00px

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/mrjones/oauth"
	"github.com/toqueteos/webbrowser"
)

var mainAPIUrl = "https://api.500px.com/v1/oauth/"

type oAuthInfo struct {
	consumerKey    string
	consumerSecret string
	OAuthToken     string
	OAuthVerifier  string
	c              *oauth.Consumer
	Port           int
}

type authResp struct {
	token    string
	verifier string
}

func (oa *oAuthInfo) Auth() {
	_, u, err := oa.c.GetRequestTokenAndUrl(fmt.Sprint("http://127.0.0.1:", oa.Port))
	if err != nil {
		log.Panicln(err)
	}

	l, err := net.Listen("tcp", fmt.Sprint(":", oa.Port))
	if err != nil {
		log.Panicln(err)
	}
	c := make(chan authResp)
	go serveOAuthResp(l, &c)

	err = webbrowser.Open(u)
	if err != nil {
		log.Panicln(err)
	}

	auth := <-c

	_ = l.Close()

	oa.OAuthToken = auth.token
	oa.OAuthVerifier = auth.verifier
}

func serveOAuthResp(l net.Listener, stop *chan authResp) {
	s := &http.Server{
		Handler: &myHandler{
			stop,
		},
	}
	_ = s.Serve(l)
}

type myHandler struct {
	a *chan authResp
}

func newOAuth(consumerKey, consumerSecret string) oAuthInfo {
	return oAuthInfo{
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
		c:              genOAuthConsumer(consumerKey, consumerSecret),
		Port:           8088,
	}
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

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	oauthToken, oauthVerifier, err := parseAccessToken(r.URL.String())
	if err != nil {
		log.Panicln(err)
	}

	fmt.Fprintln(w, "Authentication complete.")
	*h.a <- authResp{oauthToken, oauthVerifier}
}
