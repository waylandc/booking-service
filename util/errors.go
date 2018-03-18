package util

import (
	"errors"
	"net/http"
)

// Define the errors our service can encounter
var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
	ErrEmptyField      = errors.New("Request contains empty field")
)

// HTTPCodeFrom converts our error messages to an HTTP status code
func HTTPCodeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs, ErrEmptyField:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
