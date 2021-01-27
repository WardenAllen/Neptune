package util

import "strings"

func DeleteSpace(s string) string {

	return strings.Replace(s, " ", "", -1)

}

func DeleteCRLF(s string) string {

	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	return s

}
