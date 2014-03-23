package raidcheck

import (
	"io/ioutil"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("\\[([U_]+)\\]")

func checkDegraded(mdstat string) bool {
	matches := re.FindAllStringSubmatch(mdstat, -1)
	for _, match := range matches {
		if strings.Count(match[1], "U") != len(match[1]) {
			return true
		}
	}

	return false
}

func Degraded() (bool, error) {
	buf, err := ioutil.ReadFile("/proc/mdstat")
	if err != nil {
		return false, err
	}

	return checkDegraded(string(buf)), nil
}
