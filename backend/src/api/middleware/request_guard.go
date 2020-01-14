package middleware

import (
  "encoding/json"
  "log"
  "models"
  "net/http"
  "utils"
)

func RequiredAdminRole(next http.Handler) http.Handler {
  return requiredRole(isAdmin, next)
}

func RequiredGroupLeadRole(next http.Handler) http.Handler {
  return requiredRole(isGroupLead, next)
}

func RequiredGroupWorkerRole(next http.Handler) http.Handler {
  return requiredRole(isGroupWorker, next)
}

func requiredRole(isRoleMatchFn func(string)bool, next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    authToken, err := GetAuthTokenData(r)
    if err != nil {
      log.Println("error:", err)
      http.Error(w, "Forbidden", http.StatusForbidden)
      return
    }
    sessionData, found := authToken["session"]
    if !found {
      log.Println("not found session in auth token")
      http.Error(w, "Forbidden", http.StatusForbidden)
      return
    }
    var userSession models.UserSession
    err = json.Unmarshal([]byte(sessionData.(string)), &userSession)
    if err != nil {
      log.Println("error while decoding session data:", err)
      http.Error(w, "Forbidden", http.StatusForbidden)
      return
    }

    if isRoleMatchFn(userSession.Role) {
      log.Printf("authorized request %s by %s", r.URL, utils.GetFunctionName(isRoleMatchFn))
      next.ServeHTTP(w, r)
    } else {
      log.Printf("unauthorized request: %s\n", r.URL)
      http.Error(w, "Forbidden", http.StatusForbidden)
    }
  })
}
