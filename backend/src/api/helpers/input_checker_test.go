package helpers

import (
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
)

func TestIsStringCorrectTrue(t *testing.T) {
  for _, str := range correctString {
    Assert(checker.IsStringCorrect(str), func() {
      t.Logf("expected that string '%s' is correct, but it is not\n", str)
      t.Fail()
    })
  }
}

func TestIsStringCorrectFalse(t *testing.T) {
  for _, str := range incorrectString {
    Assert(!checker.IsStringCorrect(str), func() {
      t.Logf("expected that string '%s' is incorrect, but it is correct\n", str)
      t.Fail()
    })
  }
}

func TestIsLongTextCorrectTrue(t *testing.T) {
  for _, str := range correctLongText {
    Assert(checker.IsLongTextCorrect(str), func() {
      t.Logf("expected that string '%s' is correct, but it is not\n", str)
      t.Fail()
    })
  }
}

func TestIsLongTextCorrectFalse(t *testing.T) {
  for _, str := range incorrectLongText {
    Assert(!checker.IsLongTextCorrect(str), func() {
      t.Logf("expected that string '%s' is incorrect, but it is correct\n", str)
      t.Fail()
    })
  }
}
