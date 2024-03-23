package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {

    port := 8080

    fmt.Println("Welcome to simple web server")
    fmt.Println("Server running on localhost:" + strconv.Itoa(port))

    //create new serveMux for handle any route 
    mux := http.NewServeMux()

    // create handler for route: "/"
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        // add reponse header key value
        w.Header().Add("test", "add to header tested")

        //print out body or the request
        if r.Body != nil {
            buf := new(strings.Builder)
            _, err := io.Copy(buf, r.Body)
            if err != nil {
                fmt.Println((err))
            }
            fmt.Fprintf(w, buf.String())
        }
    })

    // initialize server listening for the requeste
    if err := http.ListenAndServe("localhost:" + strconv.Itoa(port), mux); err != nil {
        fmt.Println(err.Error())
    }


}
