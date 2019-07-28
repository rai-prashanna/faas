package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"log"
)


func main(){

	fmt.Println("Hello Gateway")
	RunApp()
}



func RunApp() {
	// New functionality written in Go
	log.Printf("Gateway started...")
	someFunc()

}


func someFunc() {
	factorialtarget, err := url.Parse("http://localhost:7070")
	digtarget, err := url.Parse("http://localhost:9090")

	log.Printf("forwarding to -> %s %s\n", factorialtarget.Scheme, factorialtarget.Host)
	log.Printf("forwarding to -> %s %s\n", digtarget.Scheme, digtarget.Host)


	if err != nil {
		log.Fatal(err)
	}
	factorialproxy := httputil.NewSingleHostReverseProxy(factorialtarget)
	digproxy := httputil.NewSingleHostReverseProxy(digtarget)
	http.HandleFunc("/factorial", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host // if you remove this line the request will fail... I want to debug why.

		factorialproxy.ServeHTTP(w, req)
	})
	http.HandleFunc("/dig", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host // if you remove this line the request will fail... I want to debug why.

		digproxy.ServeHTTP(w, req)
	})
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

