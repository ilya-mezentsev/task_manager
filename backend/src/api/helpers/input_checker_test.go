package helpers

import (
  "math"
  "strings"
  "testing"
  . "utils"
)

var (
  checker = NewInputChecker()
  correctString = []string{
    "hello", "whop", "1", strings.Repeat("a", 255),
  }
  incorrectString = []string{
    "", strings.Repeat("a", 256),
  }
  correctLongText= []string{
    "hello", "1", strings.Repeat("a", 1024),
  }
  incorrectLongText = []string{
    "", strings.Repeat("a", 1025),
  }
  safeUints = []uint{
    0, 10, math.MaxUint64-1,
  }
)

func TestInputChecker_IsStringCorrectTrue(t *testing.T) {
  for _, str := range correctString {
    Assert(checker.IsStringCorrect(str), func() {
      t.Logf("expected that string '%s' is correct, but it is not\n", str)
      t.Fail()
    })
  }
}

func TestInputChecker_IsStringCorrectFalse(t *testing.T) {
  for _, str := range incorrectString {
    Assert(!checker.IsStringCorrect(str), func() {
      t.Logf("expected that string '%s' is incorrect, but it is correct\n", str)
      t.Fail()
    })
  }
}

func TestInputChecker_IsLongTextCorrectTrue(t *testing.T) {
  for _, str := range correctLongText {
    Assert(checker.IsLongTextCorrect(str), func() {
      t.Logf("expected that string '%s' is correct, but it is not\n", str)
      t.Fail()
    })
  }
}

func TestInputChecker_IsLongTextCorrectFalse(t *testing.T) {
  for _, str := range incorrectLongText {
    Assert(!checker.IsLongTextCorrect(str), func() {
      t.Logf("expected that string '%s' is incorrect, but it is correct\n", str)
      t.Fail()
    })
  }
}

func TestInputChecker_IsSafeUint64True(t *testing.T) {
  for _, num := range safeUints {
    Assert(checker.IsSafeUint64(num), func() {
      t.Logf("expected that number %d is safe uint64, but it is not\n", num)
      t.Fail()
    })
  }
}

func TestInputChecker_IsSafeUint64False(t *testing.T) {
  Assert(!checker.IsSafeUint64(math.MaxUint64), func() {
    t.Fail()
  })
}
