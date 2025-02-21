package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)
const SEARCHPATH = "data/"
var directories = []string{"cheats", "sheets", "also"}

func runBat(f, lang string) ([]byte, error){
    cmd := exec.Command("bat","--color=always","--paging=never","--style=plain",
        "-l", lang,f)
    output, err := cmd.CombinedOutput()
    if err != nil {
        cmd := exec.Command("bat","--color=always","--paging=never","--style=plain",
            "-l", "bash",f)
        output2, err2 := cmd.CombinedOutput()
        if err2 != nil {
            return []byte{},err
        }
        return output2,nil
    }
    return output,nil
}
//"bat","--color=always","--paging=never","--style=plain","-l","go"
func handleReq(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    root := strings.Split(path,"/")[0]
    if path == "" {
        http.Error(w, "No command provided", http.StatusBadRequest)
        return
    }

    fmt.Println(path)


    found := false
    for _,dir := range directories {
        filePath := filepath.Join(SEARCHPATH+dir,path)
        fmt.Println(filePath)
        if output, err := runBat(filePath,root); err == nil {
            found = true
            fmt.Println(filePath)
            if dir == "also" { 
                w.Write([]byte(fmt.Sprintf("\n\033[1;33m%v\033[0m\n","See also")))
                w.Write(output)
            }else{
                w.Write([]byte(fmt.Sprintf("\033[33;1m%v:%v\033[0m\n",dir,path)))
                w.Write(output)
            }

        }
    }
    
    if !found {
        http.Error(w,http.StatusText(http.StatusTeapot) ,http.StatusTeapot)
    }


	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    //w.Write([]byte(fmt.Sprintf("\033[33;1mHello there\033[0m\n")))
}

func main() {
	http.HandleFunc("/", handleReq)
	port := "42069"
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

