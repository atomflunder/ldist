package ldist

import "testing"

func TestLower(t *testing.T) {
	s1 := "Hello, World!"
	s2 := "HELLO, WORLD!"
	ToLowercase(&s1, &s2)
	expected := "hello, world!"
	if s1 != expected || s2 != expected {
		t.Errorf("Expected both strings to be %s, but got %s and %s", expected, s1, s2)
	}
}

func TestRemoveWhitespace(t *testing.T) {
	s1 := " Hello, World! "
	s2 := " Hello, World! "
	RemoveWhitespace(&s1, &s2)
	expected := "Hello,World!"
	if s1 != expected || s2 != expected {
		t.Errorf("Expected both strings to be %s, but got %s and %s", expected, s1, s2)
	}
}

func TestRemovePunctuation(t *testing.T) {
	s1 := "Hello, World!"
	s2 := "Hello, World!"
	RemovePunctuation(&s1, &s2)
	expected := "Hello World"
	if s1 != expected || s2 != expected {
		t.Errorf("Expected both strings to be %s, but got %s and %s", expected, s1, s2)
	}
}

func TestToAlphanumeric(t *testing.T) {
	s1 := "Hello, World! 123"
	s2 := "Hello, World! 123"
	ToAlphanumeric(&s1, &s2)
	expected := "HelloWorld123"
	if s1 != expected || s2 != expected {
		t.Errorf("Expected both strings to be %s, but got %s and %s", expected, s1, s2)
	}
}
