package utils

import "regexp"

func IsValidSlug(s string) bool {
	// Define a regular expression for slug.
	re := regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

	// Check if the slug is valid.
	return re.MatchString(s)
}
