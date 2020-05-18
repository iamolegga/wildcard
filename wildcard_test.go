package wildcard

import (
	"fmt"
	"testing"
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
		if result := Match(c.value, c.pattern); result != c.expected {
			t.Errorf("%s : %s = %v", c.value, c.pattern, result)
		}
	}
}

func ExampleMatch() {
	fmt.Println(Match("foo", "f*"))
	fmt.Println(Match("bar", "b?r"))
	fmt.Println(Match("baz", "az"))
	// Output:
	// true
	// true
	// false
}
