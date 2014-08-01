package raidcheck

import (
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	failed = regexp.MustCompile("\\[([U_]+)\\]")
	resync = regexp.MustCompile("resync")
)

func CheckDegraded(mdstat string) bool {
	matches := failed.FindAllStringSubmatch(mdstat, -1)
	for _, match := range matches {
		if strings.Count(match[1], "U") != len(match[1]) {
			return true
		}
	}

	if len(resync.FindString(mdstat)) > 0 {
		return true
	}

	return false
}

func Degraded() (bool, error) {
	buf, err := ioutil.ReadFile("/proc/mdstat")
	if err != nil {
		return false, err
	}

	return CheckDegraded(string(buf)), nil
}
