package code

import (
  "fmt"
  "os"
  "testing"
  . "utils"
)

var coder Coder

func init() {
  coderKey := os.Getenv("CODER_KEY")
  if coderKey == "" {
    fmt.Println("CODER_KEY env var is not set")
    os.Exit(1)
  }

  coder = NewCoder(coderKey)
}

func TestCoder_Encrypt(t *testing.T) {
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

func TestCoder_Decrypt(t *testing.T) {
  expected := map[string]interface{}{
    "user_id": 1, "role": "admin",
  }
  decrypted, err := coder.Decrypt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyX2lkIjoxfQ.vwUFq3FTIuPbi8U6bVmzfgSHbbV5pyq6D4mrCBlvu6A")

  Assert(err == nil, func() {
    t.Log("should not be error:", err)
    t.Fail()
  })
  Assert(fmt.Sprintf("%v", decrypted["user_id"]) == fmt.Sprintf("%v", expected["user_id"]), func() {
    t.Log(GetExpectationString(expected["user_id"], decrypted["user_id"]))
    t.Fail()
  })
  Assert(decrypted["role"] == expected["role"], func() {
    t.Log(GetExpectationString(expected["role"], decrypted["role"]))
    t.Fail()
  })
}

func TestCoder_DecryptErrorEmpty(t *testing.T) {
  _, err := coder.Decrypt("")
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestCoder_DecryptErrorIncorrectFormat(t *testing.T) {
  _, err := coder.Decrypt("etJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ8.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyX2lkIjoxfQ.vwUFq3FTIuPbi8U6bVmzfgSHbbV5pyq6D4mrCBlvu6A")
  Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}
