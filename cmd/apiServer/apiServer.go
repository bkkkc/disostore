package main

import (
	"log"
	"net/http"
	"os"

	"zero3.cf/disostore/cmd/apiServer/app/heartbeat"
	"zero3.cf/disostore/cmd/apiServer/app/locate"
	"zero3.cf/disostore/cmd/apiServer/app/objects"
	"zero3.cf/disostore/cmd/apiServer/app/temp"
	"zero3.cf/disostore/cmd/apiServer/app/versions"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
