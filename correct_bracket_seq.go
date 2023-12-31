package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
			continue
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
func correctBracketSeq(seq string) bool {
	var stack []string
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' || seq[i] == '{' || seq[i] == '[' {
			stack = append(stack, string(seq[i]))
		} else {
			if len(stack) == 0 {
				return false
			}

			brck := stack[len(stack)-1]

			switch seq[i] {
			case ')':
				if brck == "(" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case ']':
				if brck == "[" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case '}':
				if brck == "{" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
