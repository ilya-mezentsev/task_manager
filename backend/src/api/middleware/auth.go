package middleware

import (
  "log"
  "net/http"
  "os"
  "plugins/code"
)

const (
  cookieAuthTokenKey = "TM-Session-Token"
  RoleAdmin = "admin"
  RoleGroupLead = "group_lead"
  RoleGroupWorker = "group_worker"
)

var coder code.Coder

func init() {
  coderKey := os.Getenv("CODER_KEY")
  if coderKey == "" {
    panic("CODER_KEY is not set")
  }

  coder = code.NewCoder(coderKey)
}

func SetAuthCookie(r *http.Request, value string) {
  r.AddCookie(CreatAuthCookie(value))
}

func GetAuthTokenData(r *http.Request) (map[string]interface{}, error) {
  cookie, err := r.Cookie(cookieAuthTokenKey)
  if err != nil {
    log.Println("auth cookie not found")
    return nil, err
  }

  decoded, err := coder.Decrypt(cookie.Value)
  if err != nil {
    log.Printf("cannot decode token: %s\n", cookie.Value)
    return nil, err
  }

  return decoded, nil
}

func isAdmin(role string) bool {
  return role == RoleAdmin
}

func isGroupLead(role string) bool {
  return role == RoleGroupLead
}

func isGroupWorker(role string) bool {
  return role == RoleGroupLead || role == RoleGroupWorker
}
