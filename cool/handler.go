package cool

import (
    "net/http"
    "fmt"
    "os"
)

func Handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, os.Getenv("cool"))
}
