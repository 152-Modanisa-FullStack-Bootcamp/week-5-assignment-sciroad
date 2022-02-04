package assignment

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {

	cases := []struct {
		a uint32
		b uint32
		r uint32
		o bool
	}{
		{1, 1, 2, false},
		{42, 2701, 2743, false},
		{42, math.MaxUint32, 41, true},
		{4294967290, 5, 4294967295, false},
		{4294967290, 10, 4, true},
		{4294967290, 6, 0, true},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v+%v", c.a, c.b), func(t *testing.T) {
			result, overflow := AddUint32(c.a, c.b)
			assert.Equal(t, c.r, result)
			assert.Equal(t, c.o, overflow)
		})
	}

}

func TestCeilNumber(t *testing.T) {

	cases := []struct {
		given float64

		result float64
	}{
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.given), func(t *testing.T) {
			result := CeilNumber(c.given)
			assert.Equal(t, c.result, result)

		})
	}

}

func TestAlphabetSoup(t *testing.T) {

	cases := []struct {
		given string

		result string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.given), func(t *testing.T) {
			result := AlphabetSoup(c.given)
			assert.Equal(t, c.result, result)

		})
	}
}

func TestStringMask(t *testing.T) {

	cases := []struct {
		given  string
		n      uint
		result string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 4, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.given), func(t *testing.T) {
			result := StringMask(c.given, c.n)
			assert.Equal(t, c.result, result)

		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"

	cases := []struct {
		given [2]string

		result string
	}{
		{[2]string{"hellocat", words}, "hello,cat"},
		{[2]string{"catbat", words}, "cat,bat"},
		{[2]string{"yellowapple", words}, "yellow,apple"},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks!", words}, "not possible"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.given), func(t *testing.T) {
			result := WordSplit(c.given)
			assert.Equal(t, c.result, result)

		})
	}

}

func TestVariadicSet(t *testing.T) {

	cases := []struct {
		given []interface{}

		result []interface{}
	}{
		{[]interface{}{4, 2, 5, 4, 2, 4}, []interface{}{4, 2, 5}},
		{[]interface{}{"bootcamp", "rocks!", "really", "rocks!"}, []interface{}{"bootcamp", "rocks!", "really"}},
		{[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.given), func(t *testing.T) {
			result := VariadicSet(c.given...)
			assert.Equal(t, c.result, result)

		})
	}

}
