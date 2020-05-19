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
		Case{"aa", "a", false},
		Case{"aa", "*", true},
		Case{"cb", "?a", false},
		Case{"adceb", "*a*b", true},
		Case{"acdcb", "a*c?b", false},
		Case{"abdbc", "a*bc", true},
		Case{"acbdbc", "a*bc", true},
		Case{"aaaabbbbcccc", "a*", true},
		Case{"a", "*aab", false},
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
