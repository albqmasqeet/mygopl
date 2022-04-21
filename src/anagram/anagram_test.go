package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   bool
	}{
		{"aba", "baa", true},
		{"aaa", "aaaa", false},
		{"aab", "abc", false},
	}
	for _, test := range tests {
		got := isAnagram(test.s1, test.s2)
		if got != test.want {
			t.Errorf("isAnagram(%q, %q), got %v, want %v", test.s1, test.s2, got, test.want)

		}
	}
}
