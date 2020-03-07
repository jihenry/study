package test

import "strings"

func addJoinStr(s []string) string {
	ret := ""
	sep := " "
	for _, arg := range s {
		ret += arg + sep
	}
	return ret
}

func joinStr(s []string) string {
	return strings.Join(s, " ")
}
