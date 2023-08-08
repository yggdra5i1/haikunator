package haikunator

import (
	"regexp"
	"testing"
)

func TestHaikunate(t *testing.T) {
	t.Run("generates a name like patient-river-3393", func(t *testing.T) {
		name := Haikunate(9999, "-")
		match, _ := regexp.MatchString(`\A\w+-\w+-\d{1,4}\z`, name)

		if !match {
			t.Errorf("Generated name %q doesn't match to pattern", name)
		}
	})

	t.Run("won't return the same name for subsequent calls", func(t *testing.T) {
		name1 := Haikunate(9999, "-")
		name2 := Haikunate(9999, "-")

		if name1 == name2 {
			t.Error("Returns the same name for subsequent calls")
		}
	})

	t.Run("permits optional configuration of the token range", func(t *testing.T) {
		name := Haikunate(9, "-")
		match, _ := regexp.MatchString(`-\d{1}\z`, name)

		if !match {
			t.Errorf("Doesn't permit optional configuration of the token range\nResult - %q", name)
		}
	})

	t.Run("drops the token if token range is 0", func(t *testing.T) {
		name := Haikunate(0, "-")
		match, _ := regexp.MatchString(`\A\w+-\w+\z`, name)

		if !match {
			t.Errorf("Doesn't drop the token if token range is 0\nResult - %q", name)
		}
	})

	t.Run("permits optional configuration of the delimiter", func(t *testing.T) {
		name := Haikunate(9999, ".")
		match, _ := regexp.MatchString(`\A\w+\.\w+\.\d{1,4}\z`, name)

		if !match {
			t.Errorf("Doesn't permit optional configuration of the delimiter\nResult - %q", name)
		}
	})

	t.Run("drops the token and delimiter if token range is 0 and delimiter empty space", func(t *testing.T) {
		name := Haikunate(0, " ")
		match, _ := regexp.MatchString(`\A\w+ \w+\z`, name)

		if !match {
			t.Errorf("Doesn't drop the token and delimiter if token range is 0 and delimiter empty space\nResult - %q", name)
		}
	})
}
