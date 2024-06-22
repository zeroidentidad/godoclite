package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/zeroidentidad/godoclite/internal"
)

func main() {
	port := flag.String("port", "8080", "Port to serve documentation")
	pkgPath := flag.String("pkg", ".", "Path to the Go package")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		internal.ServePkgDoc(w, r, *pkgPath)
	})

	log.Printf("Serving documentation on http://localhost:%s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
