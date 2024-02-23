package utils

import (
	"net/url"

	"github.com/google/uuid"
)

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func IsValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
