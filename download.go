package main

import (
	"fmt"
	"os"
	"io"
	"path/filepath"
	"net/http" 
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    fileName := filepath.Base(r.URL.Path)
    filePath := filepath.Join(storagePath, fileName)

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error retrieving file:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer file.Close()

    _, err = io.Copy(w, file)
    if err != nil {
        fmt.Println("Error downloading file:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}