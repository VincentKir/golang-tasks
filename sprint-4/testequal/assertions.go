//go:build !solution

package testequal

import (
	"reflect"
)

func supportType(kind reflect.Kind) bool {
	switch kind {
	case reflect.Slice:
		return true
	case reflect.Map:
		return true
	case reflect.String:
		return true
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	default:
		return false
	}
}

func useDeepEqual(kind reflect.Kind) bool {
	switch kind {
	case reflect.Slice:
		return true
	case reflect.Map:
		return true
	default:
		return false
	}
}

func getFormatAndArgs(head string, operator string, expected, actual interface{}, msgAndArgs ...interface{}) []interface{} {
	if len(msgAndArgs) > 0 {
		return msgAndArgs
	}
	default_format := "%s:\nexpected: %v\nactual	: %v\nmessage	: %v %s %v"
	msgAndArgs = append(
		msgAndArgs,
		default_format,
		head,
		expected,
		actual,
		expected,
		operator,
		actual,
	)
	return msgAndArgs

}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	msgAndArgs = getFormatAndArgs("equal", "!=", expected, actual, msgAndArgs...)

	if expected == nil && actual == nil {
		return true
	}

	if expected == nil || actual == nil {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		return false
	}

	expected_value := reflect.ValueOf(expected)
	actual_value := reflect.ValueOf(actual)

	expected_kind := expected_value.Kind()
	actual_kind := actual_value.Kind()

	if !supportType(expected_kind) {
		t.Errorf("%s: expected: UnSupported Type: %v", "equal", expected_value.Type())
		t.FailNow()
	}

	if !supportType(actual_kind) {
		t.Errorf("%s: actual: UnSupported Type: %v", "equal", actual_value.Type())
		t.FailNow()
	}

	if useDeepEqual(expected_kind) && useDeepEqual(actual_kind) {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
			return false
		}
		return true
	}
	if expected != actual {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
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
	msgAndArgs = getFormatAndArgs("not equal", "==", expected, actual, msgAndArgs...)

	if expected == nil && actual == nil {
		return true
	}

	if expected != nil || actual != nil {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		return false
	}

	expected_value := reflect.ValueOf(expected)
	actual_value := reflect.ValueOf(actual)

	expected_kind := expected_value.Kind()
	actual_kind := actual_value.Kind()

	if !supportType(expected_kind) {
		t.Errorf("%s: expected: UnSupported Type: %v", "not equal", expected_value.Type())
		t.FailNow()
	}

	if !supportType(actual_kind) {
		t.Errorf("%s: actual: UnSupported Type: %v", "not equal", actual_value.Type())
		t.FailNow()
	}

	if useDeepEqual(expected_kind) && useDeepEqual(actual_kind) {
		if reflect.DeepEqual(expected, actual) {
			t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
			return false
		}
		return true
	}
	if expected == actual {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		return false
	}
	return true
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	msgAndArgs = getFormatAndArgs("equal", "!=", expected, actual, msgAndArgs...)

	if expected == nil && actual == nil {
		return
	}

	if expected == nil || actual == nil {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		t.FailNow()
	}

	expected_value := reflect.ValueOf(expected)
	actual_value := reflect.ValueOf(actual)

	expected_kind := expected_value.Kind()
	actual_kind := actual_value.Kind()

	if !supportType(expected_kind) {
		t.Errorf("%s: expected: UnSupported Type: %v", "equal", expected_value.Type())
		t.FailNow()
	}

	if !supportType(actual_kind) {
		t.Errorf("%s: actual: UnSupported Type: %v", "equal", actual_value.Type())
		t.FailNow()
	}

	if useDeepEqual(expected_kind) && useDeepEqual(actual_kind) {
		if reflect.DeepEqual(expected, actual) {
			return
		}
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		t.FailNow()
	}

	if expected != actual {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		t.FailNow()
	}

}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	msgAndArgs = getFormatAndArgs("not equal", "==", expected, actual, msgAndArgs...)
	if expected == nil && actual == nil {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		t.FailNow()
	}

	if expected == nil || actual == nil {
		return
	}

	expected_value := reflect.ValueOf(expected)
	actual_value := reflect.ValueOf(actual)

	expected_kind := expected_value.Kind()
	actual_kind := actual_value.Kind()

	if !supportType(expected_kind) {
		t.Errorf("%s: expected: UnSupported Type: %v", "not equal", expected_value.Type())
		t.FailNow()
	}

	if !supportType(actual_kind) {
		t.Errorf("%s: actual: UnSupported Type: %v", "not equal", actual_value.Type())
		t.FailNow()
	}

	if useDeepEqual(expected_kind) && useDeepEqual(actual_kind) {
		if !reflect.DeepEqual(expected, actual) {
			return
		}
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		t.FailNow()
	}

	if expected == actual {
		t.Errorf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		t.FailNow()
	}
}
