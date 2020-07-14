package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		targetMux.ServeHTTP(w, r)
		log.Printf("%s %s %s %v\n", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}

func InitServer() {

	mux := http.NewServeMux()

	mux.HandleFunc("/disableUser", DisableHandler)
	mux.HandleFunc("/createUser", CreateHandler)
	mux.HandleFunc("/deleteUser", DeleteHandler)

	log.Printf("Listen On %s\n", appConfig.Listen)
	if err := http.ListenAndServe(appConfig.Listen, RequestLogger(mux)); err != nil {
		log.Println(fmt.Sprintf("listen error: %s", err))
	}

}
