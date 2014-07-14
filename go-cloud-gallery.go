package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "launchpad.net/goamz/aws"
    "launchpad.net/goamz/s3"
    "net/http"
    "os"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", rootHandler)
    r.HandleFunc("/bucket-list/{bucket}/{prefix}", bucketListHandler)
    http.Handle("/", r)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        panic(err)
    }
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to Go Cloud Gallery")
}

func bucketListHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bucketName := vars["bucket"]
    prefix := vars["prefix"]
  
    auth := aws.Auth{
        AccessKey: os.Getenv("S3_KEY"),
        SecretKey: os.Getenv("S3_SECRET"),
    }
    useast := aws.USEast

    connection := s3.New(auth, useast)
    bucket := connection.Bucket(bucketName)
    res, err := bucket.List(prefix, "", "", 1000)
    if err != nil {
        panic(err)
    }
    for _, v := range res.Contents {
        fmt.Fprint(w, bucket.URL(v.Key) + "\n")
    }
}
