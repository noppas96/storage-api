package main

import (
    "fmt"
    "net/http"
)

const storagePath = "./STORAGE"

func main() {
    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/download/", downloadHandler)
    http.HandleFunc("/delete/", deleteHandler)

    fmt.Println("Listening on :8888")
    http.ListenAndServe(":8888", nil)
}
