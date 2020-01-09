package main

import (
	"fmt"
	"net/http"
)
func handleRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my badddd site.");

}

func main() {
	// runs at all available methods.
	http.HandleFunc("/", handleRequests);
	http.ListenAndServe(":3000", nil);
}