package web

import (
	"net/http"
	"strings"

	"github.com/lamprosfasoulas/docs/pkg/files"
)

func HandleFunc(w http.ResponseWriter, r *http.Request){
    path := strings.Split(strings.TrimPrefix(r.URL.Path,"/"), "/")
    if path[0] == "" {
        http.Error(w, "No command provided", http.StatusTeapot)
        return
    }
    if strings.Contains(r.Header.Get("Accept"), "text/html") {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.Write(files.GetHTML(files.GetContent(path)))
    }else{
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.Write(files.GetContent(path))
    }
     
}
