package main

import (
	"path"
	"regexp"
	"strings"
)

func ispicture(msg string) bool {
	a := strings.Contains(msg, ".png")
	b := strings.Contains(msg, ".jpg")
	c := strings.Contains(msg, ".gif")
	d := strings.Contains(msg, ".jpeg")

	return a || b || c || d
}

func extractfile(msg string) (string, string) {
	_, filename := path.Split(msg)

	tokens := strings.Split(filename, ".")

	return tokens[0], tokens[1]
}

func extracturl(msg string) string {
	re := regexp.MustCompile("(http.*\\.(png|jpg|gif|jpeg))")

	return re.FindString(msg)
}

func isvalidhost(host string) bool {
	a, _ := regexp.MatchString("\\.[a-z]{2,}", host)

	return a
}
