package main

import (
    "errors"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    // ...

    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
    err := checkSomeThing(w, r)
    if err != nil {
        return
    }

    http.Error(w, "Operation completed!", http.StatusOK)
    fmt.Println("End of Handler.")
    return
}

func checkSomeThing(w http.ResponseWriter, r *http.Request) error{

    http.Error(w, "Bad Request!", http.StatusBadRequest)
    return errors.New("bad request")
}