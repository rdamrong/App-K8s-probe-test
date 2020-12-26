package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

var counter int = 0
func main() {
    fmt.Println("Starting, but need 30 sec to ready to serve !!")
    time.Sleep(30000 * time.Millisecond)
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":4000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	if counter < 10 {
    	  fmt.Fprintf(w, "Hello, %s :: %d :: countner = %d !\n", r.URL.Path[1:],os.Getpid(),counter)
	  counter++
        }  else {
          errorHandler(w, r, http.StatusNotFound)
        }
}



func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}
