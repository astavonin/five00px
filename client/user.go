// Package five00px provides ...
package five00px

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Sirupsen/logrus"
)

func userBy(c *http.Client, dstPoint string, vals url.Values) (*User, error) {
	log := logger.WithFields(logrus.Fields{
		"context":       "userBy",
		"lookup_method": vals,
	})

	b, err := doCommand(c, dstPoint, http.MethodGet, vals, nil)
	if err != nil {
		log.WithError(err).Warn("Failed to GET data")
		return nil, ErrUserNotFound
	}

	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(b, &objmap)

	if err != nil {
		log.WithError(err).WithField("data", string(b)).
			Error("Unable to unmarshall data")
		return nil, err
	}

	var u User
	err = json.Unmarshal(*objmap["user"], &u)
	log.WithError(err).Info("Done")

	return &u, err
}

// GetUserByID call returns User struct for a user specified by id. If id == 0
// returns the profile information for the current user.
func (f00 *Five00px) GetUserByID(id int) (*User, error) {
	dstPoint := "users"
	vals := url.Values{}
	if id != 0 {
		dstPoint += "/show"
		vals.Add("id", strconv.Itoa(id))
	}

	return userBy(f00.c, dstPoint, vals)
}

// GetUserByName returns User struct for a user specified by name.
func (f00 *Five00px) GetUserByName(name string) (*User, error) {

	return userBy(f00.c, "users/show", url.Values{"username": {name}})
}

// GetUserByEmail returns User struct for a user specified by email.
func (f00 *Five00px) GetUserByEmail(email string) (*User, error) {

	return userBy(f00.c, "users/show", url.Values{"email": {email}})
}

// ListFriends call returns list of friends for a user specified by ID.
func (f00 *Five00px) ListFriends(id int, p *Page) (*Friends, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "Friends",
		"id":      id,
		"page":    p,
	})

	b, err := doCommand(f00.c, "users/"+strconv.Itoa(id)+"/friends", http.MethodGet,
		p.Vals(), nil)

	if err != nil {
		log.WithError(err).Warn("Failed to GET data")
		return nil, err
	}

	var f Friends
	err = json.Unmarshal(b, &f)
	log.WithError(err).Info("Done")

	return &f, err
}

// ListFollowers call returns list of followers for a user specified by ID.
func (f00 *Five00px) ListFollowers(id int, p *Page) (*Followers, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "Followers",
		"id":      id,
		"page":    p,
	})
	b, err := doCommand(f00.c, "users/"+strconv.Itoa(id)+"/followers",
		http.MethodGet, p.Vals(), nil)

	if err != nil {
		log.WithError(err).Warn("Failed to GET data")
		return nil, err
	}

	var f Followers
	err = json.Unmarshal(b, &f)
	log.WithError(err).Info("Done")

	return &f, err
}

// SearchUser call returns list of users (up to one hundred) users from search
// results for a specified search term
func (f00 *Five00px) SearchUser(term string, p *Page) (*Search, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "Search",
		"term":    term,
		"page":    p,
	})
	v := p.Vals()
	v.Add("term", term)
	b, err := doCommand(f00.c, "users/search", http.MethodGet, v, nil)

	if err != nil {
		log.WithError(err).Warn("Failed to GET data")
		return nil, ErrUserNotFound
	}

	var s Search
	err = json.Unmarshal(b, &s)
	log.WithError(err).Info("Done")

	return &s, err
}

// AddFriend call adds new friend by user ID. Returns errors: ErrUserNotFound
// and ErrUserAlreadyFriend
func (f00 *Five00px) AddFriend(id int) (*User, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "DelFriend",
		"id":      id,
	})
	b, err := doCommand(f00.c, "users/"+strconv.Itoa(id)+"/friends",
		http.MethodPost, nil, nil)

	if err != nil {
		return nil, processError(log, b, errorTable{
			http.StatusNotFound:  ErrUserNotFound,
			http.StatusForbidden: ErrUserAlreadyFriend,
		})
	}

	var f User
	err = json.Unmarshal(b, &f)
	log.WithError(err).Info("Done")
	return &f, err
}

// DelFriend call deletes friend by user ID. Returns errors: ErrUserNotFound
// and ErrUserNotFriend
func (f00 *Five00px) DelFriend(id int) (*User, error) {
	log := logger.WithFields(logrus.Fields{
		"context": "DelFriend",
		"id":      id,
	})

	b, err := doCommand(f00.c, "users/"+strconv.Itoa(id)+"/friends",
		http.MethodDelete, nil, nil)

	if err != nil {
		return nil, processError(log, b, errorTable{
			http.StatusNotFound:  ErrUserNotFound,
			http.StatusForbidden: ErrUserNotFriend,
		})
	}

	var f User
	err = json.Unmarshal(b, &f)
	log.WithError(err).Info("Done")
	return &f, err
}
