package main

import (
	"log"
	"net/http"
	"os"

	"zero3.cf/disostore/cmd/dataServer/app/heartbeat"
	"zero3.cf/disostore/cmd/dataServer/app/locate"
	"zero3.cf/disostore/cmd/dataServer/app/objects"
	"zero3.cf/disostore/cmd/dataServer/app/temp"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
