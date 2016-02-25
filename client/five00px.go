// Package five00px provides main 500px API implementation
package five00px

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

// Five00px client
type Five00px struct {
	c  *http.Client
	oa oAuth
}

// New call creates and initiate Five00px object. ConsumerKey and
// ConsumerSecret have to be provided by user
func New(key, secret string) Five00px {
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

func (f00 *Five00px) Users() (string, error) {
	response, err := f00.c.Get(mainAPIUrl + "users")
	defer func() {
		_ = response.Body.Close()
	}()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(response.Body)

	return string(b), err
}

func (f00 *Five00px) Friends(id int) (string, error) {
	response, err := f00.c.Get(mainAPIUrl + "users/" + strconv.Itoa(id) + "/friends")
	defer func() {
		_ = response.Body.Close()
	}()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(response.Body)

	return string(b), err
}
