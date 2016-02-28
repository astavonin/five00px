// Package five00px provides main 500px API implementation
package five00px

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Five00px client
type Five00px struct {
	c  *http.Client
	oa oAuth
}

type Page struct {
	Rpp  int
	Page int
}

// NewPage call returns Page with 500px default values `Results Per Page` 20
// and `Page` 1
func NewPage() Page {
	return Page{20, 1}
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

func userBy(c *http.Client, dstPoint string, vals url.Values) (*User, error) {

	b, err := doGet(c, dstPoint, vals)
	if err != nil {
		return nil, ErrUserNotFound
	}

	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(b, &objmap)

	if err != nil {
		return nil, err
	}

	var u User
	err = json.Unmarshal(*objmap["user"], &u)

	return &u, err
}

// UserByID call returns User struct for a user specified by id. If id == 0
// returns the profile information for the current user.
func (f00 *Five00px) UserByID(id int) (*User, error) {
	dstPoint := "users"
	vals := url.Values{}
	if id != 0 {
		dstPoint += "/show"
		vals.Add("id", strconv.Itoa(id))
	}

	return userBy(f00.c, dstPoint, vals)
}

// UserByName returns User struct for a user specified by name.
func (f00 *Five00px) UserByName(name string) (*User, error) {

	return userBy(f00.c, "users/show", url.Values{"username": {name}})
}

// UserByEmail returns User struct for a user specified by email.
func (f00 *Five00px) UserByEmail(email string) (*User, error) {

	return userBy(f00.c, "users/show", url.Values{"email": {email}})
}

// Friends call returns list of friends for a user specified by ID.
func (f00 *Five00px) Friends(id int, page *Page) (*Friends, error) {
	vals := url.Values{}
	if page != nil {
		vals.Add("page", strconv.Itoa(page.Page))
		vals.Add("rpp", strconv.Itoa(page.Rpp))
	}
	b, err := doGet(f00.c, "users/"+strconv.Itoa(id)+"/friends", vals)

	var friends Friends
	err = json.Unmarshal(b, &friends)

	return &friends, err
}
