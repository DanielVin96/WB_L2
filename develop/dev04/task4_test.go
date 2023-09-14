package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	expected := map[string][]string{
		"газон": {"газон", "загон"},
		"кулон": {"клоун", "кулон"},
	}
	words := []string{"кулон", "клоун", "газон", "загон"}
	result := *FindAnagrams(words)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
