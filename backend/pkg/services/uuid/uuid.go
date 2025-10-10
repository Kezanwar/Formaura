package uuid

import (
	"regexp"
)

// supports UUID v1-v7
var testUUID = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-7][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

func Validate(uuid string) bool {
	return testUUID.MatchString(uuid)
}
