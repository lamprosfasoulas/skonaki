package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/lamprosfasoulas/skonaki/pkg/cache"
)
type newFile struct {
    path    string
    content string
}

func HandleAPI(w http.ResponseWriter, r *http.Request){
    // Write the api 
        if r.FormValue("path") == "" {
            http.Error(w, "Path field is required", http.StatusTeapot)
            return
        }
        if r.FormValue("content") == "" {
            http.Error(w, "Content field is required", http.StatusTeapot)
            return
        }

        var file newFile
        path := strings.Split(r.FormValue("path"),"/")
        if len(path) > 1 {
            path[0] = "_" + path[0]
        }
        file.path = filepath.Join(append([]string{"data","11.internal"},path...)...)
        file.content = r.FormValue("content")
        if err := writeAPI(file.path, &file.content); err ==nil{
            w.Write([]byte(fmt.Sprintf("File %v written successfully :)\n",file.path)))
        }else{
            w.Write([]byte(fmt.Sprintf("Writing file %v failed :(\n",file.path)))
        }
}

func writeAPI(p string, c *string) error{
    dir := filepath.Dir(p)
    if err := os.MkdirAll(dir, 0755); err != nil {
        log.Printf("%v",err)
        return err
    }else{
        file, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
        if err != nil {
            log.Println("Error opening file:", err)
            return err
        }
        defer file.Close()

        // Write the data to the file
        _, err = file.WriteString(*c)
        if err != nil {
            log.Println("Error writing to file:", err)
            return err
        }

        log.Printf("File %v written successfully",p)
        return nil
    }
}

func HandleFlush(w http.ResponseWriter, r *http.Request) {
    if e:= cache.Flush(); e != nil {
        http.Error(w, "Flush failed :(", http.StatusTeapot)
    }else{
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, "Flush successful :)")
    }
}
