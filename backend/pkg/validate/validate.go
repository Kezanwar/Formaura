package validate

import (
	"regexp"
)

func StrNotEmpty(s ...string) bool {
	for _, v := range s {
		if v == "" {
			return false
		}
	}
	return true
}

var hex_colour_regex = regexp.MustCompile(`^#(?:[0-9a-fA-F]{5})$`)

func IsHexColorCode(s ...string) bool {
	for _, v := range s {
		if !hex_colour_regex.MatchString(v) {
			return false
		}
	}
	return true
}

// Updated to support UUID v1-v7
var testUUID = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-7][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

func ValidateUUID(uuid string) bool {
	return testUUID.MatchString(uuid)
}
