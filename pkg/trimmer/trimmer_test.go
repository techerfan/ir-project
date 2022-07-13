package trimmer

import (
	"testing"
)

type testCase struct {
	word     string
	expected string
}

func TestTrim(t *testing.T) {
	testCases := []testCase{
		{
			word:     "4.5",
			expected: "4.5",
		},
		{
			word:     "word!",
			expected: "word",
		},
		{
			word:     "(word)",
			expected: "word",
		},
		{
			word:     "16/50",
			expected: "16/50",
		},
		{
			word:     "16,50",
			expected: "16,50",
		},
		{
			word:     "word.",
			expected: "word",
		},
		{
			word:     ".word",
			expected: "word",
		},
		{
			word:     "wo.rd",
			expected: "word",
		},
	}

	for _, tc := range testCases {
		trimmmed := Trim(tc.word)
		if trimmmed != tc.expected {
			t.Errorf("epected: %s got: %s\n", tc.expected, trimmmed)
		}
	}
}

func TestIsDigit(t *testing.T) {

	for _, d := range digits {
		digit := isDigit(d)
		if !digit {
			t.Fail()
		}
	}

	digit := isDigit('t')

	if digit {
		t.Fail()
	}
}

func TestIsNumber(t *testing.T) {
	isNum := isNumber("123455")
	if !isNum {
		t.Fail()
	}

	isNum = isNumber("454a5dsf")
	if isNum {
		t.Fail()
	}
}
