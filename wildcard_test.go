package wildcard_test

import (
	"fmt"
	"testing"

	"github.com/iamolegga/wildcard"
)

type Case struct {
	value    string
	pattern  string
	expected bool
}

func TestMatch(t *testing.T) {
	cases := []Case{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{"abdbc", "a*bc", true},
		{"acbdbc", "a*bc", true},
		{"aaaabbbbcccc", "a*", true},
		{"a", "*aab", false},
		{"a", "a*", true},
	}

	for _, c := range cases {
		if result := wildcard.Match(c.value, c.pattern); result != c.expected {
			t.Errorf("%s : %s = %v", c.value, c.pattern, result)
		}
	}
}

func ExampleMatch() {
	fmt.Println(wildcard.Match("foo", "f*"))
	fmt.Println(wildcard.Match("bar", "b?r"))
	fmt.Println(wildcard.Match("baz", "az"))
	// Output:
	// true
	// true
	// false
}
