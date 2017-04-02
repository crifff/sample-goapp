package app

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", helloWorld)
}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
