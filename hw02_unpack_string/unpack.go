package hw02unpackstring

import (
	"errors"
	"strconv"
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
	result := ""
	for i, ch := range arrV {
		switch {
		case i == len(arrV)-1 && unicode.IsLetter(arrV[i]):
			result += string(ch)
		case i < len(arrV)-1 && unicode.IsLetter(arrV[i]) && unicode.IsDigit(arrV[i+1]):
			result += multiRune(arrV[i], arrV[i+1])
		case i < len(arrV)-1 && unicode.IsLetter(arrV[i]) && unicode.IsLetter(arrV[i+1]):
			result += string(ch)
		case i < len(arrV)-1 && unicode.IsLetter(arrV[i]):
			continue
		}
	}
	return result
}

func multiRune(ch, c rune) string {
	result := ""
	count, _ := strconv.Atoi(string(c))
	for i := 0; i < count; i++ {
		result += string(ch)
	}
	return result
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
