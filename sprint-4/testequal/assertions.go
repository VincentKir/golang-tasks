//go:build !solution

package testequal

import (
	"reflect"
	"slices"
)

func supportType(value any) bool {
	rValue := reflect.ValueOf(value)
	switch rValue.Kind() {
	case
		reflect.Slice,
		reflect.Map,
		reflect.String,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return true
	default:
		return false
	}
}

func getFormatAndArgs(head string, operator string, expected, actual interface{}, msgAndArgs ...interface{}) (string, []interface{}) {
	if len(msgAndArgs) > 0 {
		return msgAndArgs[0].(string), msgAndArgs[1:]
	}
	defaultFormat := "%s:\nexpected: %v\nactual	: %v\nmessage	: %v %s %v"
	args := []interface{}{
		head,
		expected,
		actual,
		expected,
		operator,
		actual,
	}
	return defaultFormat, args

}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()

	msg, args := getFormatAndArgs("equal", "!=", expected, actual, msgAndArgs...)

	if !isEqual(expected, actual) {
		t.Errorf(msg, args...)
		return false
	}

	return true
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()

	msg, args := getFormatAndArgs("not equal", "==", expected, actual, msgAndArgs...)

	if isEqual(expected, actual) {
		t.Errorf(msg, args...)
		return false
	}

	return true
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	msg, args := getFormatAndArgs("equal", "!=", expected, actual, msgAndArgs...)

	if !isEqual(expected, actual) {
		t.Errorf(msg, args...)
		t.FailNow()
	}
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	msg, args := getFormatAndArgs("not equal", "==", expected, actual, msgAndArgs...)

	if isEqual(expected, actual) {
		t.Errorf(msg, args...)
		t.FailNow()
	}
}

func isEqual(expected, actual any) bool {
	if expected == nil && actual == nil {
		return true
	}

	if expected == nil || actual == nil {
		return false
	}

	if !supportType(expected) || !supportType(actual) {
		return false
	}

	if isCompareSlice(expected) && isCompareSlice(actual) {
		return сompareSlice(expected, actual)
	}

	if isCompareMap(expected) && isCompareMap(actual) {
		a := expected.(map[string]string)
		b := actual.(map[string]string)
		return compareMapStr(a, b)
	}

	return expected == actual
}

func isNilReflect(value any) bool {
	rValue := reflect.ValueOf(value)
	return rValue.IsNil()

}

func isCompareSlice(value any) bool {
	switch value.(type) {
	case []int:
		return true
	case []byte:
		return true
	default:
		return false
	}
}

func сompareSlice(a any, b any) bool {
	if isNilReflect(a) && isNilReflect(b) {
		return true
	}

	if isNilReflect(a) || isNilReflect(b) {
		return false
	}

	if isSliceInt(a) && isSliceInt(b) {
		a := a.([]int)
		b := b.([]int)
		return slices.Equal(a, b)
	}

	if isSliceByte(a) && isSliceByte(b) {
		a := a.([]byte)
		b := b.([]byte)
		return slices.Equal(a, b)
	}
	return false
}

func isSliceInt(value any) bool {
	switch value.(type) {
	case []int:
		return true
	default:
		return false
	}
}

func isSliceByte(value any) bool {
	switch value.(type) {
	case []byte:
		return true
	default:
		return false
	}
}

func isCompareMap(value any) bool {
	switch value.(type) {
	case map[string]string:
		return true
	default:
		return false
	}
}

func compareMapStr(a map[string]string, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	if isNilReflect(a) && isNilReflect(b) {
		return true
	}

	if isNilReflect(a) || isNilReflect(b) {
		return false
	}

	for kA, vA := range a {
		if vB, ok := b[kA]; !ok || vB != vA {
			return false
		}
	}
	return true
}
