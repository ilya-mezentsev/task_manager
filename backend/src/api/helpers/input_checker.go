package helpers

import "math"

type InputChecker struct {
  minStringLength, maxStringLength, maxLongStringLength uint
}

func NewInputChecker() InputChecker {
  return InputChecker{
    minStringLength: 0,
    maxStringLength: 255,
    maxLongStringLength: 1024,
  }
}

func (ic InputChecker) IsSafeUint64(num uint) bool {
  return num >= 0 && num < math.MaxUint64
}

func (ic InputChecker) IsStringCorrect(str string) bool {
  stringLength := uint(len([]byte(str)))
  return ic.minStringLength < stringLength && stringLength <= ic.maxStringLength
}

func (ic InputChecker) IsLongTextCorrect(text string) bool {
  stringLength := uint(len([]byte(text)))
  return ic.minStringLength < stringLength && stringLength <= ic.maxLongStringLength
}
