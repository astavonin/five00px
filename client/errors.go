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

// ErrInvalidInput ...
var ErrInvalidInput = errors.New("Invalid user input")

// ErrPhotoNotFound ...
var ErrPhotoNotFound = errors.New("Photo with the specified ID does not exist")

// ErrPhotoNotAvailable
var ErrPhotoNotAvailable = errors.New("The photo was either deleted, belongs to a deactivated user")

var ErrVoteRejected = errors.New("The vote has been rejected; common reasons are:" +
	" current user is inactive, has not completed their profile, is trying to vote" +
	" on their own photo, or has already voted for the photo")

var ErrBadComment = errors.New("The body of the comment was not specified")

var ErrUnprocessableEntity = errors.New("The system had trouble saving the record. You may retry again.")

type five00Error struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
