package main

import (
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	result, err := process("data.txt")
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]int{
		"alice":     1,
		"cried":     1,
		"and":       1,
		"curiouser": 2,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v but got %v", expected, result)
	}
}
