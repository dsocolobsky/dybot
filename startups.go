package main

import (
	"strings"
	log "github.com/cihub/seelog"
	"os/exec"
)

func startups() string {
	a, _ := exec.Command("shuf", "-n", "1", "startups.txt").Output()
	//cmd.Stdout = &a

	log.Tracef("a: %s", a)
	
	b, _ := exec.Command("shuf", "-n 1", "dict.txt").Output()
	//cmd.Stdout = &b

	log.Tracef("b: %s", b)

	ab := strings.Trim(string(a[:]), "\n")
	ac := strings.Trim(string(b[:]), "\n")
	ad := strings.Trim(ab + " but for " +  ac, "\n")

	return ad
}
