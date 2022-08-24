package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/time",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", time.Now().Format("2.Jan.2006 15:04:05"))
		})
	http.ListenAndServe("localhost:8080", nil)

}
