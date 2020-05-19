package wildcard_test

import (
	"fmt"
	"testing"

	"github.com/iamolegga/wildcard"
)

func TestMatch(t *testing.T) {
	type testCase struct {
		value    string
		pattern  string
		expected bool
	}

	cases := []testCase{
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
		t.Run(fmt.Sprintf("value: %s, pattern: %s", c.value, c.pattern), func(t *testing.T) {
			if result := wildcard.Match(c.value, c.pattern); result != c.expected {
				t.Errorf("expected: %v, got: %v", c.expected, result)
			}
		})
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

func TestIsPattern(t *testing.T) {
	type testCase struct {
		value    string
		expected bool
	}

	cases := []testCase{
		{"foo", false},
		{"*", true},
		{"?", true},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("value: %s", c.value), func(t *testing.T) {
			if result := wildcard.IsPattern(c.value); result != c.expected {
				t.Errorf("expected: %v, got: %v", c.expected, result)
			}
		})
	}
}

func ExampleIsPattern() {
	fmt.Println(wildcard.IsPattern("f*"))
	fmt.Println(wildcard.IsPattern("b?r"))
	fmt.Println(wildcard.IsPattern("baz"))
	// Output:
	// true
	// true
	// false
}
