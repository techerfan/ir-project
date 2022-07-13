package trimmer

import "strings"

var digits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var punctuations = []rune{'.', ',', '$', '%', '!', '?', '(', ')', '/', '-', '_', '"', '\'', ':', ';'}

// Trim removes punctuations.
func Trim(word string) string {
	trimmedWord := ""
	isNum := true
	for i, r := range word {
		isPuntuation := false
		for _, p := range punctuations {
			if r == p && (p == '.' || p == '/' || p == ',') && isNum {
				isRestNum := isNumber(word[i+1:])
				if isRestNum {
					return word
				} else {
					isPuntuation = true
				}
			} else if r == p {
				isPuntuation = true
			}
		}
		if !isPuntuation {
			trimmedWord = trimmedWord + string(r)
		}
		if !isDigit(r) {
			isNum = false
		}
	}
	return strings.ToLower(trimmedWord)
}

// isDigit checks if the rune is digit or not.
func isDigit(r rune) bool {
	for _, digit := range digits {
		if r == digit {
			return true
		}
	}
	return false
}

// isNumber ensures that the text is number.
func isNumber(text string) (is_number bool) {
	is_number = true
	for _, r := range text {
		digit := isDigit(r)
		if !digit {
			is_number = false
			return
		}
	}
	return
}
