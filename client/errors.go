// Package five00px provides ...
package five00px

import "errors"

// ErrUserNotFound error.
var ErrInternal = errors.New("Internal error")
var ErrUserNotFound = errors.New("User does not exist in our database")
var ErrUserAlreadyFriend = errors.New("The user requested has been disabled or already in followers list")

type five00Error struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
