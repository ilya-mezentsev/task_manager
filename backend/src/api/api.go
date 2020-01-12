package api

import (
  "api/middleware"
  "encoding/json"
  "github.com/gorilla/mux"
  "io/ioutil"
  "log"
  "net/http"
)

var router *mux.Router

func init() {
  router = mux.NewRouter()
  router.Use(middleware.RequiredAuthCookieOrHeader)
}

type ErrorResponse struct {
  Status string `json:"status"`
  ErrorDetail string `json:"error_detail"`
}

type SuccessResponse struct {
  Status string `json:"status"`
  Data interface{} `json:"data"`
}

func decodeRequestBody(r *http.Request, target interface{}) {
  requestBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(ReadRequestBodyError)
  }

  err = json.Unmarshal(requestBody, target)
  if err != nil {
    panic(CannotDecodeRequestBody)
  }
}

func encodeAndSendResponse(w http.ResponseWriter, v interface{}) {
  output, _ := json.Marshal(SuccessResponse{
    Status: "ok",
    Data: v,
  })

  w.Header().Set("content-type", "application/json")
  if _, err := w.Write(output); err != nil {
    log.Println("Error while trying to write response:", err)
    panic(CannotWriteResponse)
  }
}

func sendErrorIfPanicked(w http.ResponseWriter) {
  if err := recover(); err != nil {
    log.Println("Panicked:", err)

    output, _ := json.Marshal(ErrorResponse{
      Status: "error",
      ErrorDetail: err.(error).Error(),
    })

    w.Header().Set("content-type", "application/json")
    if _, err = w.Write(output); err != nil {
      log.Println("Error while trying to write response:", err)
      http.Error(w, "internal server error", http.StatusInternalServerError)
    }
  }
}
