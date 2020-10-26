//
// This is free and unencumbered software released into the public domain.
// For more information, please refer to <https://unlicense.org>
//

package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// Setup the flags
	port := flag.String("p", "8080", "host files on port")
	directory := flag.String("d", "./", "host files in directory")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	// Get the IP address
	var ipAddress net.IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Printf("ERROR: failed to get IP address")
	} else {
		defer conn.Close()
		ipAddress = conn.LocalAddr().(*net.UDPAddr).IP
	}

	// Serve the files
	log.Printf("Hosting: %s on: %s:%s", *directory, ipAddress, *port)
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatalf("%s failed to host files: %s\n", os.Args[0], err)
	}
}
