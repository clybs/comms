package cmd

import (
	"strings"
)

func getCommandParams(v string) []string {
	p := normalizeInput(v)
	return strings.Split(p, " ")
}

func normalizeInput(v string) string {
	v = strings.Join(strings.Fields(v), " ")
	return strings.TrimSpace(v)
}
