package util

import "strings"

func ParseStringToList(str string, sep string) []string {
	var res []string
	for _, s := range strings.Split(str, sep) {
		res = append(res, strings.TrimSpace(s))
	}
	return res
}
