package greetings


import (
	"testing"
	"regexp"
)


func TestHelloName(t *testing.T) {

	name := "Adrian"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Adrian") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestEmptyHello(t *testing.T) {

	msg,err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") =%q, %v, want "", error`, msg,err)
	}
}
