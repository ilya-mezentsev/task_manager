package middleware

import (
  "log"
  "net/http"
  "plugins/code"
)

const (
  tokenKey = "TM-Session-Token"
  RoleAdmin = "admin"
  RoleGroupLead = "group_lead"
  RoleGroupWorker = "group_worker"
)

var coder code.Coder

func init() {
  coder = code.NewCoder("123456789012345678901234")
}

func SetTokenForAdmin(r *http.Request) {
  r.Header.Set(tokenKey, getTokenWithRole(RoleAdmin))
}

func SetTokenForGroupLead(r *http.Request) {
  r.Header.Set(tokenKey, getTokenWithRole(RoleGroupLead))
}

func SetTokenForGroupWorker(r *http.Request) {
  r.Header.Set(tokenKey, getTokenWithRole(RoleGroupWorker))
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
  token := r.Header.Get(tokenKey)
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
