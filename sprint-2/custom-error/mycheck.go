//go:build !solution

package mycheck

import (
	"errors"
	"strings"
	"unicode"
)

var FoundNumbers = errors.New("found numbers")
var LineIsTooLong = errors.New("line is too long")
var NoTwoSpaces = errors.New("no two spaces")

type MyError []error

func (m MyError) Error() string {
	str := make([]string, len(m))

	for i, err := range m {
		str[i] = err.Error()
	}
	return strings.Join(str, ";")
}

func MyCheck(input string) error {
	var err MyError

	if hasNumber(input) {
		err = append(err, FoundNumbers)
	}

	if isLongLine(input) {
		err = append(err, LineIsTooLong)
	}

	if !hasExactlyTwoSpaces(input) {
		err = append(err, NoTwoSpaces)
	}

	if err != nil {
		return err
	}

	return nil
}

func hasNumber(input string) bool {
	for _, i := range input {
		if unicode.IsDigit(i) {
			return true
		}
	}
	return false
}

func isLongLine(input string) bool {
	if len(input) > 20 {
		return true
	}
	return false
}

func hasExactlyTwoSpaces(input string) bool {
	spaces := 0
	for _, i := range input {
		if unicode.IsSpace(i) {
			spaces++
		}
		if spaces > 2 {
			return false
		}
	}
	if spaces < 2 {
		return false
	}
	return true
}
