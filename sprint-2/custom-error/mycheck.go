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

type MyError struct {
	errors []error
}

func (m MyError) Error() string {
	var str []string
	for _, err := range m.errors {
		str = append(str, err.Error())
	}
	return strings.Join(str, ";")
}

func MyCheck(input string) error {
	err := MyError{}

	if checkErrOfNumber(input) {
		err.errors = append(err.errors, FoundNumbers)
	}

	if checkErrOfLongLine(input) {
		err.errors = append(err.errors, LineIsTooLong)
	}

	if checkErrOfTwoSpaces(input) {
		err.errors = append(err.errors, NoTwoSpaces)
	}

	if err.errors != nil {
		return err
	}

	return nil
}

func checkErrOfNumber(input string) bool {
	for _, i := range input {
		if unicode.IsDigit(i) {
			return true
		}
	}
	return false
}

func checkErrOfLongLine(input string) bool {
	if len(input) > 20 {
		return true
	}
	return false
}

func checkErrOfTwoSpaces(input string) bool {
	spaces := 0
	for _, i := range input {
		if unicode.IsSpace(i) {
			spaces++
		}
	}
	if spaces == 2 {
		return false
	}
	return true
}
