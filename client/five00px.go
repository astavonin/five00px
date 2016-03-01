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

// Page structure stores RPP(Results Per Page) and Page(Return the specified
// page of the resource) values.
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
func (f00 *Five00px) Friends(id int, p *Page) (*Friends, error) {
	b, err := doGet(f00.c, "users/"+strconv.Itoa(id)+"/friends", pageToVals(p))

	if err != nil {
		return nil, err
	}

	var f Friends
	err = json.Unmarshal(b, &f)

	return &f, err
}

// Followers call returns list of followers for a user specified by ID.
func (f00 *Five00px) Followers(id int, p *Page) (*Followers, error) {
	b, err := doGet(f00.c, "users/"+strconv.Itoa(id)+"/followers", pageToVals(p))

	if err != nil {
		return nil, err
	}

	var f Followers
	err = json.Unmarshal(b, &f)

	return &f, err
}

// Search call returns list of users (up to one hundred) users from search
// results for a specified search term
func (f00 *Five00px) Search(term string, p *Page) (*Search, error) {
	v := pageToVals(p)
	v.Add("term", term)
	b, err := doGet(f00.c, "users/search", v)

	if err != nil {
		return nil, err
	}

	var s Search
	err = json.Unmarshal(b, &s)

	return &s, err
}

func (f00 *Five00px) AddFriend(id int) (*User, error) {
	b, err := doPost(f00.c, "users/"+strconv.Itoa(id)+"/friends")

	if err != nil {
		var e00 five00Error
		err = json.Unmarshal(b, &e00)
		if err != nil {
			return nil, ErrInternal
		}
		switch e00.Status {
		case http.StatusNotFound:
			return nil, ErrUserNotFound
		case http.StatusForbidden:
			return nil, ErrUserAlreadyFriend
		}
	}

	var f User
	err = json.Unmarshal(b, &f)
	return &f, err
}
