package datastructures

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet()
	if s == nil {
		t.Fail()
	}
}

func TestAddContainsExisting(t *testing.T) {
	s := NewSet()
	s.Add("foo")

	if !s.Contains("foo") {
		t.Fail()
	}
}

func TestContainsNonexisting(t *testing.T) {
	s := NewSet()

	if s.Contains("foo") {
		t.Fail()
	}
}

func TestLength(t *testing.T) {
	s := NewSet()
	s.Add("foo")
	s.Add("bar")
	s.Add("baz")

	if s.Length() != 3 {
		t.Fail()
	}
}

func TestLengthWhenAddingDuplicates(t *testing.T) {
	s := NewSet()
	s.Add("foo")
	s.Add("bar")
	s.Add("baz")
	s.Add("baz")

	if s.Length() != 3 {
		t.Fail()
	}
}

func TestIntegers(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(2)

	if s.Length() != 2 {
		t.Fail()
	}

	if !s.Contains(1) {
		t.Fail()
	}
}

func TestIterate(t *testing.T) {
	s := NewSet()
	s.Add("foo")
	s.Add("bar")
	s.Add("baz")

	for v := range s.Iterate() {
		fmt.Println(v)
	}
}

func TestUnion(t *testing.T) {
	s1 := NewSet()
	s2 := NewSet()

	s1.Add("foo")
	s1.Add("bar")

	s2.Add("baz")
	s2.Add("boo")

	s3 := s1.Union(s2)
	if s3.Length() != 4 {
		t.Fail()
	}
}
