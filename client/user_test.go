// Package five00px provides ...
package five00px

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"testing"

	"github.com/Sirupsen/logrus"
)

func handleDelFriend(id string) (string, int) {
	switch id {
	case "42":
		return "friends_del.json", http.StatusOK
	case "100":
		return "friends_no_friend.json", http.StatusForbidden
	}
	return "404.json", http.StatusNotFound
}

func handleAddFriend(id string) (string, int) {
	switch id {
	case "42":
		return "friends_already.json", http.StatusForbidden
	case "100":
		return "friends_add.json", http.StatusOK
	}
	return "404.json", http.StatusNotFound
}

func handleSearch(u *url.URL) (string, int) {
	m, _ := url.ParseQuery(u.RawQuery)
	if m["term"][0] == "@@@" {
		return "search_empty.json", http.StatusOK
	}
	return "search.json", http.StatusOK
}

func handleShow(u *url.URL) (string, int) {
	if u.RawQuery == "" {
		return "400.json", http.StatusBadRequest
	}
	m, _ := url.ParseQuery(u.RawQuery)
	for k, v := range m {
		switch k {
		case "username":
			if v[0] == "alexanderstavonin" {
				return "user.json", http.StatusOK
			}
		case "id":
			if v[0] == "9091479" {
				return "user.json", http.StatusOK
			}
		case "email":
			if v[0] == "alex@sysdev.me" {
				return "user.json", http.StatusOK
			}
		}
	}
	return "404.json", http.StatusNotFound
}

func userHandler(w http.ResponseWriter, r *http.Request) (string, int) {

	u, err := url.Parse(r.URL.String())
	if err != nil {
		logrus.Error(err)
		return "", http.StatusInternalServerError
	}

	reFriends := regexp.MustCompile(`/users/(\w+)/friends`)
	reFollowers := regexp.MustCompile(`/users/(\w+)/followers`)

	if u.Path == "/users" {
		return "user.json", http.StatusOK
	} else if u.Path == "/users/search" {
		return handleSearch(u)
	} else if u.Path == "/users/show" {
		return handleShow(u)
	} else if res := reFriends.FindStringSubmatch(u.Path); len(res) > 0 && r.Method == http.MethodDelete {
		return handleDelFriend(res[1])
	} else if res := reFriends.FindStringSubmatch(u.Path); len(res) > 0 && r.Method == http.MethodPost {
		return handleAddFriend(res[1])
	} else if res := reFriends.FindStringSubmatch(u.Path); len(res) > 0 {
		if res[1] == "9091479" {
			return "friends.json", http.StatusOK
		}
	} else if res := reFollowers.FindStringSubmatch(u.Path); len(res) > 0 {
		if res[1] == "9091479" {
			return "followers.json", http.StatusOK
		}
	}

	return "", http.StatusInternalServerError
}

func TestGetUserByName_Fail(t *testing.T) {
	f00 := NewTest500px()

	_, err := f00.UserByName("bad_name")
	if err == nil {
		t.Fatal("It's fail case")
	}
}

func TestGetUserByEmail(t *testing.T) {
	f00 := NewTest500px()

	u, err := f00.UserByEmail("alex@sysdev.me")
	if err != nil {
		t.Fatal(err)
	}
	if u.Username != "alexanderstavonin" || u.ID != 9091479 {
		t.Fatal("Invalid user data")
	}
}

func TestGetUserByName(t *testing.T) {
	f00 := NewTest500px()

	u, err := f00.UserByName("alexanderstavonin")
	if err != nil {
		t.Fatal(err)
	}
	if u.Username != "alexanderstavonin" || u.ID != 9091479 {
		t.Fatal("Invalid user data")
	}
}

func TestGetUserByID(t *testing.T) {
	f00 := NewTest500px()

	u, err := f00.UserByID(0)
	if err != nil {
		t.Fatal(err)
	}
	if u.Username != "alexanderstavonin" || u.ID != 9091479 {
		t.Fatal("Invalid user data")
	}
	_, err = f00.UserByID(9091479)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFriends(t *testing.T) {
	f00 := NewTest500px()

	page := NewPage()
	u, err := f00.Friends(9091479, &page)
	if err != nil {
		t.Fatal(err)
	}
	if u.FriendsCount != 25 || u.FriendsPages != 2 || len(u.Users) != 20 {
		t.Fatal("Invalid user data")
	}
}

func TestFollowers(t *testing.T) {
	f00 := NewTest500px()

	page := NewPage()
	u, err := f00.Followers(9091479, &page)
	if err != nil {
		t.Fatal(err)
	}
	if u.FollowersCount != 15 || u.FollowersPages != 1 || len(u.Users) != 15 {
		t.Fatal("Invalid user data")
	}
}

func TestSearch(t *testing.T) {
	f00 := NewTest500px()

	page := NewPage()
	s, err := f00.Search("@@@", &page) // will not find
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("QQQQ", s.TotalItems, s.CurrentPage, len(s.Users))
	if s.TotalItems != 0 || s.CurrentPage != 1 || len(s.Users) != 0 {
		t.Fatal("Invalid user data")
	}
	// ----

	page.Rpp = 19
	s, err = f00.Search("empty", &page)
	if err != nil {
		t.Fatal(err)
	}
	if s.TotalItems != 7715 || s.CurrentPage != 1 || len(s.Users) != 19 {
		t.Fatal("Invalid user data")
	}
}

func TestAddFriend(t *testing.T) {
	f00 := NewTest500px()

	_, err := f00.AddFriend(42) // already friends
	if err != ErrUserAlreadyFriend {
		t.Fatal("Should be error here")
	}

	_, err = f00.AddFriend(0) // not exists
	if err != ErrUserNotFound {
		t.Fatal("Should be error here")
	}

	_, err = f00.AddFriend(100) // new friend
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelFriend(t *testing.T) {
	f00 := NewTest500px()

	_, err := f00.DelFriend(42) // we are friends
	if err != nil {
		t.Error(err)
	}

	_, err = f00.DelFriend(0) // not exists
	if err != ErrUserNotFound {
		t.Errorf("Expecting \"%s\", found \"%s\"", ErrUserNotFound, err)
	}

	_, err = f00.DelFriend(100) // not a friend
	if err != ErrUserNotFriend {
		t.Errorf("Expecting \"%s\", found \"%s\"", ErrUserNotFriend, err)
	}
}
