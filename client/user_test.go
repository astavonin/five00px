// Package five00px provides ...
package five00px

import (
	"net/http"
	"net/url"
	"regexp"
	"testing"

	"github.com/Sirupsen/logrus"
)

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

	if u.Path == "/users" {
		return "user.json", http.StatusOK
	} else if u.Path == "/users/show" {
		return handleShow(u)
	} else if res := reFriends.FindStringSubmatch(u.Path); len(res) > 0 {
		if res[1] == "9091479" {
			return "friends.json", http.StatusOK
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

	u, err := f00.Friends(9091479)
	if err != nil {
		t.Fatal(err)
	}
	if u.FriendsCount != 25 || u.FriendsPages != 2 || len(u.Users) != 20 {
		t.Fatal("Invalid user data")
	}
}
