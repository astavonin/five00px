// Package five00px provides ...
package five00px

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/Sirupsen/logrus"
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

	log := logger.WithFields(logrus.Fields{
		"context": "Auth",
	})
	reqToken, u, err := oa.c.GetRequestTokenAndUrl(fmt.Sprint("http://127.0.0.1:", oa.Port))
	if err != nil {
		log.WithError(err).Error("Unable to create OAuth provider")
		return nil, ErrAuth
	}

	l, err := net.Listen("tcp", fmt.Sprint(":", oa.Port))
	if err != nil {
		log.WithError(err).Error("Cannot start TCP service")
		return nil, ErrAuth
	}
	c := make(chan authResp)
	go serveOAuthResp(l, &c)

	err = webbrowser.Open(u)
	if err != nil {
		log.WithError(err).Error("Cannot start browser")
		return nil, ErrAuth
	}

	auth := <-c

	_ = l.Close()

	if auth.err != nil {
		log.WithError(auth.err).Error("Authentication error")
		return nil, auth.err
	}
	accessToken, err := oa.c.AuthorizeToken(reqToken, auth.verifier)
	if err != nil {
		log.WithError(err).Error("Unable to authorize token")
		return nil, err
	}

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
	tokens := val["oauth_token"]
	verifiers := val["oauth_verifier"]
	if tokens != nil && verifiers != nil {
		oauthToken = tokens[0]
		oauthVerifier = verifiers[0]
	}

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
