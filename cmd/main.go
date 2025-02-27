package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/lamprosfasoulas/skonaki/pkg/cache"
	"github.com/lamprosfasoulas/skonaki/pkg/files"
	"github.com/lamprosfasoulas/skonaki/pkg/web"
)
// These are the env variables used

// SKON_REDIS_ADDR
// SKON_REDIS_PASSWD
// SKON_ALLOW_API
// SKON_ALLOW_SUGGEST
// SKON_DOMAIN


func main() {
    // Initialise the redis cache
    cache.InitRedis()
    cmd := exec.Command("ls","data")
    if dirs, err := cmd.CombinedOutput(); err == nil{
        files.DIRS = strings.Fields(string(dirs))
    }

    if os.Getenv("SKON_DOMAIN") == ""{
        web.Domain = "localhost:42069"
    }else{
        web.Domain = os.Getenv("SKON_DOMAIN")
    }
    // Init the web server
	http.HandleFunc("/", web.HandleFunc)
    if os.Getenv("SKON_ALLOW_SUGGEST") == "true" {
        http.HandleFunc("/:suggest", web.HandleSug)
    }
    if os.Getenv("SKON_ALLOW_API") == "true" {
        http.HandleFunc("/:api", web.HandleAPI)
    }
	port := "42069"
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

