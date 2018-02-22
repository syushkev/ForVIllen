package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindVersion123(t *testing.T) {
	t.Run("Positive_not_default_RE", func(t *testing.T) {
		testSlice, err := FindVersion123("5.6.9.3.9", `(\d+\.\d+)\.(.+)`)
		assert.Equal(t, []string{"5.6", "9.3.9"}, testSlice)
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}
	})
	
	t.Run("Positive_not_default_RE_2", func(t *testing.T) {
		testSlice, err := FindVersion123("5.6.9.3.9", `()(\d.+)`)
		assert.Equal(t, []string{"", "5.6.9.3.9"}, testSlice)
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}
	})

	t.Run("Positive_default_RE", func(t *testing.T) {
		testSlice, err := FindVersion123("5.6.9.3.9", `(\d+)\.(.+)`)
		assert.Equal(t, []string{"5", "6.9.3.9"}, testSlice)
		if err != nil {
			t.Fatalf("unexpected error %s", err)
		}
	})

	t.Run("Empty_Version", func(t *testing.T) {
		_, err := FindVersion123("", `(\d+)\.(.+)`)
		if err == nil {
			t.Fatalf("unexpected Nil")
		}
	})

	t.Run("Invalid_Version", func(t *testing.T) {
		_, err := FindVersion123("sadasd", `(\d+)\.(.+)`)
		if err == nil {
			t.Fatalf("unexpected Nil")
		}
	})

	t.Run("Invalid_Expression", func(t *testing.T) {
		_, err := FindVersion123("5.6.9.3.9", `/(.)(.)\`)
		if err == nil {
			t.Fatalf("unexpected Nil")
		}
	})
}
