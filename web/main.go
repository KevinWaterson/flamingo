package main

import (
  "log"
  "net/http"
  "time"
)

func logHandler(next http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%v: %v", r.Method, r.URL.Path)
    next.ServeHTTP(w, r)
  }
}

func main() {
  http.Handle("/est", logHandler(http.HandlerFunc(est)))

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func est(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().In(time.FixedZone("EST", 0)).Format(time.RFC1123)
  w.Write([]byte("The time is: " + tm))
}
