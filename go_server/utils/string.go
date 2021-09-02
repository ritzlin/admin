package utils

import "strings"

func IsWhiteSpace(str string) bool {
	return strings.TrimSpace(str) == ""
}