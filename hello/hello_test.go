package hello

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGoodName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg := Good("Gladys")
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q,  want match for %#q`, msg, want)
	}
}
