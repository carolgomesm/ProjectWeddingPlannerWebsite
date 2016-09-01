package main

import (
    "net/http"
    "log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
		log.Fatal(err)
    }

	log.Printf("Server running on 9000. Open your browser on http://localhost:9000")
    err = http.ListenAndServe(":9000", http.FileServer(http.Dir(dir)))
    if err != nil {
        log.Printf("Error running web server for static assets: %v", err)
    }
}
