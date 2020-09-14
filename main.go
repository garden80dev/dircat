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
		endpoint := fmt.Sprintf("/%s/", key)
		http.Handle(
			endpoint,
			http.StripPrefix(endpoint, http.FileServer(http.Dir(dir.Path))),
		)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
