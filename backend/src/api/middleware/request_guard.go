package middleware

import (
  "log"
  "net/http"
  "plugins/code"
)

const (
  tokenKey = "TM-Session-Token"
  roleAdmin = "admin"
)

var coder = code.NewCoder("123456789012345678901234")

func RequiredAdminRole(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get(tokenKey)
    decoded, err := coder.Decrypt(token)
    if err != nil {
      log.Printf("cannot decode token: %s\n", token)
      log.Println("error:", err)
      http.Error(w, "Forbidden", http.StatusForbidden)
    }

    if role, found := decoded["role"]; found && role == roleAdmin {
      next.ServeHTTP(w, r)
    } else {
      log.Printf("unauthorized request: %s\n", r.URL)
      http.Error(w, "Forbidden", http.StatusForbidden)
    }
  })
}
