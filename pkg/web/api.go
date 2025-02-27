package web

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)
type newFile struct {
    path    string
    content string
}

func HandleAPI(w http.ResponseWriter, r *http.Request){
    // Write the api 
    if r.Method == http.MethodPost{
        var file newFile
        path := strings.Split(r.FormValue("path"),"/")
        if len(path) > 1 {
            path[0] = "_" + path[0]
        }
        file.path = filepath.Join(append([]string{"data","11.internal"},path...)...)
        file.content = r.FormValue("content")
        writeAPI(file.path, &file.content)
    }else{
        HandleFunc(w,r)
    }
}

func writeAPI(p string, c *string) {
    dir := filepath.Dir(p)
    if err := os.MkdirAll(dir, 0755); err != nil {
        log.Printf("%v",err)
	}else{
    file, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return 
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.WriteString(*c)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}

	log.Printf("File %v written successfully",p)
    }
}
