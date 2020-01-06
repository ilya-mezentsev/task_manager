package utils

import "fmt"

func Assert(condition bool, onFalseFn func()) {
  if !condition {
    onFalseFn()
  }
}

func AssertErrorsEqual(err1, err2 error, onFalseFn func()) {
  Assert(err1 != nil && err2 != nil && err1.Error() == err2.Error(), onFalseFn)
}

func GetExpectationString(expected, got interface{}) string {
  return fmt.Sprintf("expected: %v, got: %v\n", expected, got)
}
