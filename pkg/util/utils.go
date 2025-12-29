package util

import "strings"

func BuildCode(name string) string {
	return strings.ToUpper(strings.ReplaceAll(name, " ", ""))
}

func BuildAlias(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
