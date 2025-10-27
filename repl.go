package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	clean := strings.Fields(lower)
	return clean
}
