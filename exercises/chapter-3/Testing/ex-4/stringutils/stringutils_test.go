package stringutils

import "testing"

func FuzzIsPalindrome(f *testing.F) {
	f.Add("racecar")
	f.Add("hello")
	f.Add("")

	f.Fuzz(func(t *testing.T, s string) {
		if IsPalindrome(s) {
			if !IsPalindrome(Reverse(s)) {
				t.Errorf("Reverse of palindrome %q is not palindrome: %q", s, Reverse(s))
			}
		}
	})
}
