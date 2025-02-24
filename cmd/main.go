package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/lamprosfasoulas/skonaki/pkg/cache"
	"github.com/lamprosfasoulas/skonaki/pkg/files"
	"github.com/lamprosfasoulas/skonaki/pkg/web"
)


func main() {
    // Initialise the redis cache
    cache.InitRedis()
    cmd := exec.Command("ls","data")
    if dirs, err := cmd.CombinedOutput(); err == nil{
        files.DIRS = strings.Fields(string(dirs))
    }

    // Init the web server
	http.HandleFunc("/", web.HandleFunc)
	port := "42069"
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

