package haikunator

import (
	"regexp"
	"testing"
)

func TestHaikunate(t *testing.T) {
	name := Haikunate(9999, "-")
	match, _ := regexp.MatchString(`\A\w+-\w+-\d{1,4}\z`, name)

	if !match {
		t.Errorf("Generated name %q doesn't match to pattern", name)
	}
}
