// Package five00px provides main 500px API implementation
package five00px

import (
	"encoding/json"
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

func doGet(c *http.Client, dstPoint string) ([]byte, error) {
	response, err := c.Get(mainAPIUrl + dstPoint)
	defer func() {
		_ = response.Body.Close()
	}()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (f00 *Five00px) User() (*User, error) {
	b, err := doGet(f00.c, "users")

	var u User
	err = json.Unmarshal(b, u)

	return &u, err
}

func (f00 *Five00px) Friends(id int) (*Friends, error) {
	b, err := doGet(f00.c, "users/"+strconv.Itoa(id)+"/friends")

	var friends Friends

	err = json.Unmarshal(b, &friends)

	return &friends, err
}
