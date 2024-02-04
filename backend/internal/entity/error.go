package entity

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrPostNotFound = errors.New("post not found")
var ErrCommentNotFound = errors.New("comment not found")
var ErrInvalidLogin = errors.New("login shouldn't be empty")
var ErrInvalidPassword = errors.New("len of password should be > 8")
var ErrInvalidAuthHeader = errors.New("invalid auth header")
var ErrEmptyAuthHeader = errors.New("empty auth header")
