package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Lucaaaa Zelmans!")
	fmt.Println("Ca bien marché!!!")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8080", nil)

}
