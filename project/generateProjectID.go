package project

import (
	"math/rand"
	"net/url"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randHash return a random hash of n size
func randHash(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// encodeProjectName return url standard project name
// Trim the project, eliminate special charaters
// and then add "-" in between white spaces,
func encodeProjectName(name string) string {

	name = strings.Trim(name, " ")
	name = strings.ReplaceAll(name, " ", "-") + "-"
	name = url.PathEscape(name)

	return name
}

// GenerateID return a unique project id
func GenerateID(name string) string {
	endHash := randHash(10)
	name = encodeProjectName(name) + endHash
	return name
}
