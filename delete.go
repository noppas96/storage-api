package main

import(
	"fmt"
	"os"
	"path/filepath"
	"net/http" 
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    fileName := filepath.Base(r.URL.Path)
    filePath := filepath.Join(storagePath, fileName)

    err := os.Remove(filePath)
    if err != nil {
        fmt.Println("Error deleting file:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}