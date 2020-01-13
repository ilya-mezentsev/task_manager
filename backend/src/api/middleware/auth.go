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

func setAuthCookie(r *http.Request, value string) {
  r.AddCookie(CreatAuthCookie(value))
}

func SetTokenForAdmin(r *http.Request) {
  setAuthCookie(r, RoleAdmin)
}

func SetTokenForGroupLead(r *http.Request) {
  setAuthCookie(r, RoleGroupLead)
}

func SetTokenForGroupWorker(r *http.Request) {
  setAuthCookie(r, RoleGroupWorker)
}

func getAuthTokenData(r *http.Request) (map[string]interface{}, error) {
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

func isAdmin(tokenData map[string]interface{}) bool {
  role, found := tokenData["role"]
  return found && role == RoleAdmin
}

func isGroupLead(tokenData map[string]interface{}) bool {
  role, found := tokenData["role"]
  return found && role == RoleGroupLead
}

func isGroupWorker(tokenData map[string]interface{}) bool {
  role, found := tokenData["role"]
  // coz group lead is group worker too
  return found && (role == RoleGroupLead || role == RoleGroupWorker)
}
