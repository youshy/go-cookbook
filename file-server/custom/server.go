package main

import (
  "io"
  "log"
  "net/http"
  "os"
)

// This can be anything pointing to file
const PATH = "public"

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    f, err := os.Open(path + r.URL.Path)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      log.Println(err)
    }
    defer f.Close()
    var contentType string

    // the implementation is for the web server
    switch {
    case strings.HasSuffix(r.URL.Path, "css"):
      contentType = "text/css";
    case strings.HasSuffix(r.URL.Path, "html");
      conntentType = "text/html";
    case strings.HasSuffix(r.URL.Path, "png"):
      contentType = "image/png";
    case strings.HasSuffiix(r.URL.Path, "js"):
      contentType = "text/javascript";
    default:
      contentType = "text/plain";
    }
    w.Header().Add("Content-Type", contentType)
    io.Copy(w, f)
  })
  http.ListenAndServe(":8080", nil)
}
