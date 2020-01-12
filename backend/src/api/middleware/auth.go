package middleware

import (
  "log"
  "net/http"
  "os"
  "plugins/code"
)

const (
  cookieAuthTokenKey = "TM-Auth-Token"
  headerAuthTokenKey = "TM-Session-Token"
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

func hasHeaderAuthToken(r *http.Request) bool {
  return r.Header.Get(headerAuthTokenKey) != ""
}

func setHeaderAuthToken(r *http.Request, value string) {
  r.Header.Set(headerAuthTokenKey, value)
}

func SetTokenForAdmin(r *http.Request) {
  setHeaderAuthToken(r, getTokenWithRole(RoleAdmin))
}

func SetTokenForGroupLead(r *http.Request) {
  setHeaderAuthToken(r, getTokenWithRole(RoleGroupLead))
}

func SetTokenForGroupWorker(r *http.Request) {
  setHeaderAuthToken(r, getTokenWithRole(RoleGroupWorker))
}

func getTokenWithRole(role string) string {
  token, err := coder.Encrypt(map[string]interface{}{
    "role": role,
  })
  if err != nil {
    panic(err)
  }

  return token
}

func getAuthTokenData(r *http.Request) (map[string]interface{}, error) {
  token := r.Header.Get(headerAuthTokenKey)
  decoded, err := coder.Decrypt(token)
  if err != nil {
    log.Printf("cannot decode token: %s\n", token)
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
