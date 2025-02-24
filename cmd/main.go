package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lamprosfasoulas/skonaki/pkg/cache"
	"github.com/lamprosfasoulas/skonaki/pkg/web"
)


func main() {
    // Initialise the redis cache
    cache.InitRedis()

    // Init the web server
	http.HandleFunc("/", web.HandleFunc)
	port := "42069"
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

