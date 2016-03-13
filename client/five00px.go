// Package five00px provides main 500px API implementation
package five00px

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Five00px client
type Five00px struct {
	c  *http.Client
	oa oAuth
}

var logger *logrus.Logger = nil

// New call creates and initiate Five00px object. ConsumerKey and
// ConsumerSecret have to be provided by user
func New(key, secret string, log *logrus.Logger) Five00px {
	if log == nil {
		logger = logrus.New()
		logger.Level = logrus.ErrorLevel
	} else {
		logger = log
	}
	return Five00px{
		oa: newOAuth(key, secret),
	}
}

// Auth initiate OAuth authentication call. Default Web broser will be
// popped up during authentication. Store AccessToken for futher usage with
// Restore API call. Returns error on authorization failure.
func (f00 *Five00px) Auth() (*AccessToken, error) {
	t, err := f00.oa.Auth()
	if err != nil {
		return nil, err
	}
	f00.c, err = f00.oa.createClient(t)
	return t, nil
}

// Restore call restores OAuth session without additional authentication call.
// Does not require to show any additional requests.
func (f00 *Five00px) Restore(t *AccessToken) error {
	c, err := f00.oa.createClient(t)
	if err != nil {
		return err
	}
	f00.c = c
	return nil
}
