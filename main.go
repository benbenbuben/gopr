package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

const dataFile = "data.txt"

type Data struct {
    String string `json:"string"`
}

func main() {
    http.HandleFunc("/write", writeHandler)
    http.HandleFunc("/read", readHandler)
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Printf("Server listening on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
    str := r.URL.Query().Get("string")
    if str == "" {
        http.Error(w, "Missing 'string' query parameter", http.StatusBadRequest)
        return
    }
    
    err := ioutil.WriteFile(dataFile, []byte(str), 0644)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
    content, err := ioutil.ReadFile(dataFile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    data := Data{String: string(content)}
    json.NewEncoder(w).Encode(data)
}