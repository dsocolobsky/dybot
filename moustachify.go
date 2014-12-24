package main

import (
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

func extracturl(msg string) string {
	re := regexp.MustCompile("(http.*\\.(png|jpg|gif|jpeg))")

	return re.FindString(msg)
}

func moustachify(url string) string {
	return "http://mustachify.me/?src=" + extracturl(url)
}
