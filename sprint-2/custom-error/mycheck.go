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

	hasErrNum, hasErrNoTwoSpace := checkNumAndSpaces(input)

	if hasErrNum {
		err = append(err, FoundNumbers)
	}

	if isLongLine(input) {
		err = append(err, LineIsTooLong)
	}

	if hasErrNoTwoSpace {
		err = append(err, NoTwoSpaces)
	}

	if err != nil {
		return err
	}

	return nil
}

func isLongLine(input string) bool {
	if len(input) > 20 {
		return true
	}
	return false
}

func checkNumAndSpaces(input string) (bool, bool) {
	spaces := 0
	err_space := false
	err_num := false
	for _, i := range input {
		if spaces > 2 {
			err_space = true
		}

		if err_num && err_space {
			return err_num, err_space
		}

		if unicode.IsSpace(i) {
			spaces++
			continue
		}

		if unicode.IsDigit(i) {
			err_num = true
		}

	}
	if spaces < 2 {
		err_space = true
	}
	return err_num, err_space
}
