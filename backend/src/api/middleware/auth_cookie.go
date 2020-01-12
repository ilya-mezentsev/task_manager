package middleware

import (
  "net/http"
  "time"
)

const (
  authCookieKey = "TM-Auth-Token"
)

func CreatAuthCookie(value string) *http.Cookie {
  token, err := coder.Encrypt(map[string]interface{}{
    "role": value,
  })
  if err != nil {
    panic(err)
  }

  return &http.Cookie{
    Name: authCookieKey,
    Value: token,
    Path: "/",
    HttpOnly: true,
    MaxAge: 3600,
  }
}

func GetExpiredAuthCookie() *http.Cookie {
  return &http.Cookie{
    Name: authCookieKey,
    Value: "",
    Path: "/",
    Expires: time.Unix(0, 0),
  }
}
