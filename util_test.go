package newstring

import (
	"strings"
	"testing"
)

type _M map[string]string

const (
	separator = " ¶ "
)

func runTestCases(t *testing.T, converter func(string) string, cases map[string]string) {
	for k, v := range cases {
		s := converter(k)

		if s != v {
			t.Fatalf("case fails. [case:%v]\nshould => %#v\nactual => %#v", k, v, s)
		}
	}
}

func sep(strs ...string) string {
	return strings.Join(strs, separator)
}

func split(str string) []string {
	return strings.Split(str, separator)
}
