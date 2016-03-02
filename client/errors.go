// Package five00px provides ...
package five00px

import "errors"

// ErrInternal ...
var ErrInternal = errors.New("Internal error")

// ErrUserNotFound ...
var ErrUserNotFound = errors.New("User does not exist in our database")

// ErrUserAlreadyFriend ...
var ErrUserAlreadyFriend = errors.New("The user requested has been disabled or already in followers list")

// ErrUserNotFriend ...
var ErrUserNotFriend = errors.New("The user requested has been disabled or not in followers list")

type five00Error struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
