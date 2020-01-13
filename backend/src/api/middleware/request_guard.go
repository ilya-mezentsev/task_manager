package middleware

import (
  "log"
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

func requiredRole(isRoleMatchFn func(map[string]interface{})bool, next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    tokenData, err := getAuthTokenData(r)
    if err != nil {
      log.Println("error:", err)
      http.Error(w, "Forbidden", http.StatusForbidden)
    }

    if isRoleMatchFn(tokenData) {
      log.Printf("authorized request %s by %s", r.URL, utils.GetFunctionName(isRoleMatchFn))
      next.ServeHTTP(w, r)
    } else {
      log.Printf("unauthorized request: %s\n", r.URL)
      http.Error(w, "Forbidden", http.StatusForbidden)
    }
  })
}
