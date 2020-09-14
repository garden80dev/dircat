package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	cfg, err := readConfig(filepath.Join(HOME, ".dircat.toml"))
	if err != nil {
		log.Fatal(err)
	}

	for key, dir := range cfg.Dirs {
		endpoint := fmt.Sprintf("/%s", key)
		http.Handle(
			endpoint,
			http.StripPrefix(endpoint, http.FileServer(http.Dir(dir.Path))),
		)
	}

	// Simple static webserver:
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
	// http.Handle("/sas", http.FileServer(http.Dir("H:\\WorkSpaceNaoto\\OTO_VIEWER\\viewer\\src\\assets\\images")))
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("H:\\WorkSpaceNaoto\\OTO_VIEWER\\viewer\\src\\assets\\images"))))
}
