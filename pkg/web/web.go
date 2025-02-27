package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/lamprosfasoulas/skonaki/pkg/files"
)

type Page struct{
    Domain      string
    Path        string
    List        string
    Content     template.HTML
}

var Domain string 

func isTerminal(userAgent string) bool{
    progs := []string{"curl","wget","HTTPie","fetch"}
    for _,prog := range progs {
        if strings.Contains(userAgent, prog){
            return true
        }
    }
    return false
}

//func returnResponse(r *http.Request) ([]byte, string) {
//}

func HandleFunc(w http.ResponseWriter, r *http.Request){
    var p Page
    p.Domain = Domain
    if isTerminal(r.Header.Get("User-Agent")) {
        //return for terminal
        path := strings.Split(strings.TrimPrefix(r.URL.Path,"/"), "/")
        start := time.Now()
        response := files.GetContent(path)
        log.Printf("Terminal request %v with response time: %v\n",path,time.Since(start))
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.Write(*response)
        //return response, "text/plain; charset=utf-8"
    }else{
        //return for browsers
        tmpl, _ := template.ParseFiles("html/index.html")
        path := strings.Split(r.URL.Query().Get("path"),"/")
        if r.URL.Query().Get("path") == "" {
            path = strings.Split(strings.TrimPrefix(r.URL.Path,"/"), "/")
        }
        start := time.Now()
        //outBat := files.GetContent(path)
        //var response bytes.Buffer
        //response = files.GetHTML(outBat)
        p.Content = template.HTML(files.GetHTML(files.GetContent(path)))
        fmt.Printf("%v",p.Content)
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        tmpl.Execute(w,p)
        log.Printf("HTML request %v  with response time: %v\n",path,time.Since(start))
        //return response.Bytes(), "text/html; charset=utf-8"
    }
    //c, t := returnResponse(r)
    //w.Header().Set("Content-Type", t)//"text/html; charset=utf-8"
    //w.Write(c)
}
