package main

import "fmt"

func main() {
	TestPalindrome()
}

func TestPalindrome() {
	testCases := []struct {
		name     string
		s        string
		expected bool
		desc     string
	}{
		{"basic palindrome", "A man, a plan, a canal: Panama", true, "Classic palindrome with spaces and punctuation"},
		{"not palindrome", "race a car", false, "Not a palindrome"},
		{"empty string", "", true, "Empty string is palindrome"},
		{"single char", "a", true, "Single character is palindrome"},
		{"single non-alpha", ".", true, "Single non-alphanumeric becomes empty, so palindrome"},
		{"alphanumeric palindrome", "Madam", true, "Simple word palindrome"},
		{"mixed case palindrome", "RaceCar", true, "Mixed case palindrome"},
		{"numbers palindrome", "12321", true, "Numeric palindrome"},
		{"alphanumeric mixed", "A1B2b1a", true, "Mix of letters and numbers"},
		{"not palindrome mixed", "A1B2c1a", false, "Mix that's not palindrome"},
		{"only punctuation", ".,!@#", true, "Only non-alphanumeric chars"},
		{"palindrome with numbers", "Was it a car or a cat I saw?", false, "Tricky case"},
		{"actual car palindrome", "Madam, I'm Adam", true, "Another classic"},
		{"numbers and letters", "12a21", true, "Simple mixed palindrome"},
		{"spaces only", "   ", true, "Only spaces"},
	}

	for _, tc := range testCases {
		result := isPalindrome(tc.s)

		fmt.Printf("Test: %s - ", tc.desc)

		if result == tc.expected {
			fmt.Println("PASS")
		} else {
			fmt.Printf("FAIL (expected %t, got %t)\n", tc.expected, result)
		}
	}
}

func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}

		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}

		left++
		right--
	}

	return true

}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLowerCase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}

	return c
}
