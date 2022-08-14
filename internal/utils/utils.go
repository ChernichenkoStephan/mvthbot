package utils

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/exp/utf8string"
)

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

func GenPassword(length int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rand.Seed(time.Now().UnixNano())
	valid := utf8string.NewString(s)
	var (
		min = 0
		max = len(s) - 1
	)

	buffer := make([]rune, length)
	for i := 0; i < length; i++ {
		buffer[i] = valid.At(rand.Intn(max-min) + min)
	}

	return string(buffer)
}
