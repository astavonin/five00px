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

// AccessToken is an alias for oauth.AccessToken structure
type AccessToken oauth.AccessToken

type oAuth struct {
	c    *oauth.Consumer
	t    *oauth.AccessToken
	Port int
}

type authResp struct {
	token    string
	verifier string
	err      error
}

func (oa *oAuth) Auth() (*AccessToken, error) {
	reqToken, u, err := oa.c.GetRequestTokenAndUrl(fmt.Sprint("http://127.0.0.1:", oa.Port))
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

	if auth.err != nil {
		return nil, auth.err
	}
	accessToken, err := oa.c.AuthorizeToken(reqToken, auth.verifier)

	token := AccessToken(*accessToken)
	return &token, nil
}

func (oa *oAuth) createClient(t *AccessToken) (*http.Client, error) {
	at := oauth.AccessToken(*t)
	return oa.c.MakeHttpClient(&at)
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

func newOAuth(key, secret string) oAuth {
	return oAuth{
		c:    genOAuthConsumer(key, secret),
		Port: 8088,
	}
}

func genOAuthConsumer(consumerKey, consumerSecret string) *oauth.Consumer {
	return oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   mainAPIUrl + "oauth/request_token",
			AuthorizeTokenUrl: mainAPIUrl + "oauth/authorize",
			AccessTokenUrl:    mainAPIUrl + "oauth/access_token",
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
		*h.a <- authResp{"", "", err}
	}

	fmt.Fprintln(w, "Authentication complete.")
	*h.a <- authResp{oauthToken, oauthVerifier, nil}
}
