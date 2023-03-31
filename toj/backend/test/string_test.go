package test

import "testing"
import "strings"

func TestString(t *testing.T) {
	s := "123\n"

	println(s)

	println(strings.Count(s, "\n"))

	s = strings.TrimRight(s, "\n")
	println(s)
	println(strings.Count(s, "\n"))
	



}
