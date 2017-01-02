package choose

import (
	"reflect"
	"testing"
)

type chooseStringInput struct {
	s string
	n int
}

var chooseTests = []struct {
	given chooseStringInput
	expected []string
} {
	{chooseStringInput{"ab", 2}, []string{"ab"}},
	{chooseStringInput{"ab", 1}, []string{"a", "b"}},
	{chooseStringInput{"ab", 6}, []string{"ab"}},
	{chooseStringInput{"abc", 2}, []string{"ab", "ac", "bc"}},
	{chooseStringInput{"abc", 1}, []string{"a", "b", "c"}},
	{chooseStringInput{"abc", 3}, []string{"abc"}},
	{chooseStringInput{"abcd", 3}, []string{"abc", "abd", "acd", "bcd"}},	
	{chooseStringInput{"abcd", 2}, []string{"ab", "ac", "bc", "ad", "bd", "cd"}},	
	{chooseStringInput{"日本語", 2}, []string{"日本", "日語", "本語"}},	
	{chooseStringInput{"ab", 0}, []string{}},
}

func TestChooseString(t *testing.T) {
	for _, tt := range chooseTests {
		actual := ChooseString(tt.given.s, tt.given.n)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Choose(\"%s\", %d): expected %v, got %v\n", tt.given.s, tt.given.n, tt.expected, actual)
		}
	}	
}

func TestLarge(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	k := 4
	out := ChooseString(s, k)
	if (len(out) != 14950) {
		t.Errorf("Choose(\"%s\", %d): expected 14950, got %d\n", s, k, len(out))
	}
}