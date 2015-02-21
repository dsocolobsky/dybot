package main

import (
	"errors"
	log "github.com/cihub/seelog"
	"net"
	"strings"
)

func ping(host string) bool {
	log.Debugf("Received %s", host)

	if strings.Contains(host, "https://") {
		log.Debug("Has https")
		host = strings.Trim(host, "https://")
	}
	if strings.Contains(host, "http://") {
		log.Debug("Has http")
		host = strings.Trim(host, "http://")
	}
	if strings.Contains(host, "www.") {
		log.Debug("Has www")
		host = strings.Trim(host, "www.")
	}

	host = host + ":80"

	log.Infof("Pinging host %s", host)

	_, err := net.Dial("tcp", host)
	if err != nil {
		return false
	}

	return true
}

func isup(host string) (bool, error) {
	if !isvalidhost(host) {
		return false, errors.New("Invalid host")
	}

	return ping(host), nil
}
