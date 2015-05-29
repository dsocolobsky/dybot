package main

import (
	"net/http"
	"strconv"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
	log "github.com/cihub/seelog"
)

var months = map[string]int {
	"january": 1, "february": 2, "march": 3, "april": 4,
	"may": 5, "june": 6, "july": 7, "august": 8, "september": 9,
	"october": 10, "november": 11, "december": 12,
}

var responses = [4]string {
	"Did you know that ",
	"Let me tell you that ",
	"You know, one curious thing about ",
	"I've heard that ",
}

func randomresponse() string {
	rand.Seed(time.Now().UTC().UnixNano())
	response := responses[rand.Intn(len(responses))]
	return response
}

func promptyear(msg string) string {
	year := getyear(msg)
	data := runyearquery(year)
	response := randomresponse() + data

	return response
}

func promptdate(msg string) string {
	month := getmonth(msg)
	day := getday(msg)

	data := rundatequery(month, day)
	response := randomresponse() + data

	return response
}

func runyearquery(year string) string {
	query := "http://numbersapi.com/" + year + "/year"
	res, _ := http.Get(query)
	content, _ := ioutil.ReadAll(res.Body)
	return string(content)
}

func rundatequery(month string, day string) string {
	log.Debug(month)
	log.Debug(day)

	return "BOGUS"

	query := "http//numbersapi.com/" + month + "/" + day + "/date"
	res, _ := http.Get(query)
	content, _ := ioutil.ReadAll(res.Body)
	return string(content)
}


func isyear(msg string) bool {
	val := getyear(msg)
	if val != "" {
		return true
	}

	return false
}

// Check if there's a date in a message (e.g. 15 of May)
func isdate(msg string) bool {
	month := getmonth(msg)
	if month == "" {
		return false
	}

	day := getday(msg)
	if day == "" {
		return false
	}

	return true
}

// Extract a year as string given a message
func getyear(msg string) string {
	tokens := strings.Split(msg, " ")

	for _, token := range tokens {
		if len(string(token)) == 4 && string(token[0]) == "1" || string(token[0]) == "2" {
			return token
		}
	}

	return ""
}

// Get the month requested given the message
func getmonth(msg string) string {
	sane := strings.ToLower(msg)
	sane = strings.TrimSpace(sane)

	for key, _ := range months {
		if strings.Contains(sane, key) {
			return strconv.Itoa(months[key])
		}
	}

	return ""
}

// Get the day requested given the message
func getday(msg string) string {
	sane := strings.ToLower(msg)
	sane = strings.TrimSpace(sane)

	for i := 1; i <= 12; i++ {
		if strings.Contains(sane, strconv.Itoa(i)) {
			return strconv.Itoa(i)
		}
	}

	return ""
}
