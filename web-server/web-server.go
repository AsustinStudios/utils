package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	server := os.Args[1]
	file, err := os.OpenFile("/logs/web-server", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Printf("Web server for '%s' started.\n", server)
	err = http.ListenAndServe(":80", http.FileServer(http.Dir("/web/"+server)))
	if err != nil {
		log.Println(err)
		log.Fatalf("Web server for '%s' terminated.\n", server)
	}

	log.Println("Web server for '%s' terminated.\n", server)
}
