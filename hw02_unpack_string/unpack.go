package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(v string) (string, error) {
	arrV := []rune(v)
	switch {
	case len(arrV) == 0:
		return "", nil
	case !validate(arrV):
		return "", ErrInvalidString
	default:
		return buildString(arrV), nil
	}
}

func buildString(arrV []rune) string {
	var builder strings.Builder
	for i, ch := range arrV {
		switch {
		case i == len(arrV)-1 && unicode.IsLetter(arrV[i]):
			builder.WriteString(string(ch))
		case i < len(arrV)-1 && unicode.IsLetter(arrV[i]) && unicode.IsDigit(arrV[i+1]):
			count, _ := strconv.Atoi(string(arrV[i+1]))
			builder.WriteString(strings.Repeat(string(arrV[i]), count))
		case i < len(arrV)-1 && unicode.IsLetter(arrV[i]) && unicode.IsLetter(arrV[i+1]):
			builder.WriteString(string(ch))
		case i < len(arrV)-1 && unicode.IsLetter(arrV[i]):
			continue
		}
	}
	return builder.String()
}

func validate(v []rune) bool {
	if (len(v) == 1 && unicode.IsDigit(v[0])) || unicode.IsDigit(v[0]) {
		return false
	}
	for i := range v {
		if i < len(v)-1 && unicode.IsDigit(v[i]) && unicode.IsDigit(v[i+1]) {
			return false
		}
	}
	return true
}
