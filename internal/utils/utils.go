package utils

import "strings"

var runeMap map[string]string = map[string]string{
	"%21": "!",
	"%23": "#",
	"%24": "$",
	"%26": "&",
	"%27": "'",
	"%28": "(",
	"%29": ")",
	"%2A": "*",
	"%2B": "+",
	"%2C": ",",
	"%2F": "/",
	"%3A": ":",
	"%3B": ";",
	"%3D": "=",
	"%3F": "?",
	"%40": "@",
	"%5B": "[",
	"%5D": "]",
}

func DecodeLF(line string) string {
	res := line
	for code, char := range runeMap {
		res = strings.ReplaceAll(res, code, char)
	}
	return res
}

func Pop[T any](stack []T) (T, []T) {
	op := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return op, stack
}

func GenPassword() string {
	return "password"
}
