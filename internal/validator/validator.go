package validator

import (
	"net/url"
	"strings"
)

type Validator struct {
	FieldErrors map[string]string
}

// NotBlank() returns true if a value is not an empty string.
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// ValidURL() returns true if the value has a valid URL format.
func ValidURL(value string) bool {
	_, err := url.Parse(value)
	return err == nil
}
