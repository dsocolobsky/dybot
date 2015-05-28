package main

import (
	log "github.com/cihub/seelog"
	"github.com/koyachi/go-nude"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

func randomstring() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func downloadimage(rawurl string) (string, error) {
	res, err := http.Get(rawurl)
	if err != nil {
		log.Debug("Can't GET the url")
		return "", err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug("Can't read the data")
		return "", err
	}

	res.Body.Close()

	filename, extension := extractfile(rawurl)
	log.Debug(filename)
	log.Debug(extension)

	finalname := filename + "." + extension

	ioutil.WriteFile(finalname, data, 0666)
	return finalname, nil
}

func isnsfw(rawurl string) (bool, error) {
	filename, err := downloadimage(rawurl)
	if err != nil {
		log.Debug("downloadimage failed")
		return false, err
	}

	nsfw, err := nude.IsNude(filename)
	if err != nil {
		log.Debug("Can't load image")
		return false, err
	}

	err = os.Remove(filename)
	if err != nil {
		log.Debug("Failure during removal of image")
	}

	return nsfw, nil
}
