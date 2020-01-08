package code

import (
  "testing"
  . "utils"
)

var coder = NewCoder("123456789012345678901234")

func TestEncrypt(t *testing.T) {
  expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyX2lkIjoxfQ.vwUFq3FTIuPbi8U6bVmzfgSHbbV5pyq6D4mrCBlvu6A"
  encrypted, err := coder.Encrypt(map[string]interface{}{
    "user_id": 1, "role": "admin",
  })

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(encrypted == expected, func() {
    t.Log(GetExpectationString(expected, encrypted))
    t.Fail()
  })
}

func TestDecrypt(t *testing.T) {
  expected := map[string]interface{}{
    "user_id": 1, "role": "admin",
  }
  decrypted, err := coder.Decrypt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyX2lkIjoxfQ.vwUFq3FTIuPbi8U6bVmzfgSHbbV5pyq6D4mrCBlvu6A")

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(len(decrypted) == len(expected), func() {
    t.Log(GetExpectationString(expected, decrypted))
    t.Fail()
  })
}

func TestDecryptErrorEmpty(t *testing.T) {
  _, err := coder.Decrypt("")
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestDecryptErrorIncorrectFormat(t *testing.T) {
  _, err := coder.Decrypt("etJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ8.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyX2lkIjoxfQ.vwUFq3FTIuPbi8U6bVmzfgSHbbV5pyq6D4mrCBlvu6A")
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}
