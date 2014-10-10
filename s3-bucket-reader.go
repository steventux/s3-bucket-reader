package main

import (
    "encoding/json"
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
    r.HandleFunc("/list/{bucket}/{prefix}", listHandler)
    http.Handle("/", r)

    err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
    
    if err != nil {
        panic(err)
    }
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to Go Cloud API Reader")
}

var bucketContents = func(r *http.Request) map[string]string {
    urlMap := make(map[string]string)

    bucketName := mux.Vars(r)["bucket"]
    prefix := mux.Vars(r)["prefix"]

    auth := aws.Auth{
        AccessKey: os.Getenv("S3_KEY"),
        SecretKey: os.Getenv("S3_SECRET"),
    }
    region := aws.Regions[os.Getenv("S3_REGION")]

    connection := s3.New(auth, region)
    bucket := connection.Bucket(bucketName)

    res, err := bucket.List(prefix, "", "", 1000)

    if err != nil {
        panic(err)
    }
    for _, v := range res.Contents {
        urlMap[v.Key] = bucket.URL(v.Key)
    }

    return urlMap
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    urlMap := bucketContents(r)
    urlJson, err := json.Marshal(urlMap)

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(urlJson)
}
