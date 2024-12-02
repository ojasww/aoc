package main

import (
	"reflect"
	"testing"
)

/*
*
14   16
86   77
49   92
11   76
33   51
*/
func Test_ExtractInput(t *testing.T) {

	first := []int{}
	second := []int{}

	input := Input{
		first,
		second,
	}

	input.ExtractInput("test_input.txt")

	expectd_first := []int{14, 86, 49, 11, 33}
	expectd_second := []int{16, 77, 92, 76, 51}

	if !reflect.DeepEqual(first, expectd_first) || !reflect.DeepEqual(second, expectd_second) {
		t.Fatalf("Inputs are not Equal")
	}
}
