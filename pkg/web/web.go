package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/lamprosfasoulas/docs/pkg/files"
)
func isTerminal(userAgent string) bool{
    progs := []string{"curl","wget","HTTPie","fetch"}
    for _,prog := range progs {
        if strings.Contains(userAgent, prog){
            return true
        }
    }
    return false
}

func HandleFunc(w http.ResponseWriter, r *http.Request){
    if isTerminal(r.Header.Get("User-Agent")) {
        path := strings.Split(strings.TrimPrefix(r.URL.Path,"/"), "/")
        if path[0] == "" {
            //landing page
            path[0] = "home"
        }
        if path[0] == ":list" {
            fmt.Println("command")
            w.Write(files.GetList())
            return
        }
        outBat := files.GetContent(path)
        //return for terminal
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.Write(outBat)
    }else{
        path := strings.Split(r.URL.Query().Get("path"),"/")
        if path[0] == "" {
            //landing page
            path[0] = "home"
        }
        outBat := files.GetContent(path)
        //return for browsers
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        tmpl, _ := template.ParseFiles("html/index.html")
        tmpl.Execute(w,template.HTML(files.GetHTML(outBat)))
        //w.Write(files.GetHTML(outBat))
    }
    fmt.Print()
     
}
