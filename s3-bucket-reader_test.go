package main

import(
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestListHander(t *testing.T) {
    w := httptest.NewRecorder()

    req, err := http.NewRequest("GET", "http://localhost/list/foo/bar", nil)

    bucketContents = func(r *http.Request) map[string]string {
      contents := make(map[string]string)
      contents ["foo"] = "http://www.foo.com"
      contents ["bar"] = "http://www.bar.com"
      return contents
    }

    listHandler(w, req)

    if err != nil {
        fmt.Printf(w.Body.String())
    }

    contentTypeHeader := w.Header()["Content-Type"][0]
    if contentTypeHeader != "application/json" {
        t.Error("Expected 'application/json' Content-Type header, got ", contentTypeHeader)
    }

}
