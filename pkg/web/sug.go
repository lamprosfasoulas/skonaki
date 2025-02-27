package web

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//type page struct{
//    Domain  string
//    Path    string
//    List    string
//    Content string
//}

func HandleSug(w http.ResponseWriter, r *http.Request){
    var p Page
    tmpl, _ := template.ParseFiles("html/sugform.html")
    p.Domain = Domain 
    if isTerminal(r.Header.Get("User-Agent")){
        HandleFunc(w, r)
        return
    }
    path := r.URL.Query().Get("file")
    if path != "" {
        //find file append it to textview and then overwrite
        //tmpl, _ := template.ParseFiles("html/sugform.html")
        if c,e:= getSug(path);e == nil{
            p.Content = template.HTML(c)
        }else{
            p.Content = template.HTML(" ")
        }
        p.Path = path
        //tmpl.Execute(w,p)
    }
        //tmpl, _ := template.ParseFiles("html/sugform.html")
        //tmpl.Execute(w,p)
    if r.Method == http.MethodPost{
        path := r.FormValue("path")
        c := r.FormValue("sug-text")
        writeSug(path,&c)
    }
    p.List = getList()
    tmpl.Execute(w,p)
}

func getList() string{
    //contList := "The suggestions we already have are:\n\n"
    var contList string
    err := filepath.Walk("suggestions", func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            log.Println(err)
            return nil
        }
        if !info.IsDir(){
            dispPath := strings.Split(path,"/")[1:]
            if len(dispPath) > 1 {
                dispPath[0] = strings.Split(dispPath[0],"_")[1]
            }
            contList += fmt.Sprintf("curl skonaki.it.auth.gr/%v\n",filepath.Join(dispPath...))
        }
        return nil
    })
    if err != nil {
        return "" 
    }
    return contList
}

func getSug(p string) ([]byte, error){
    p = "_" + p
    fmt.Println("this is get Sup",filepath.Join("suggestions",p))
    content, err := os.ReadFile(filepath.Join("suggestions",p))
    if err != nil {
        log.Printf("Error reading file: %v",err)
        return nil, err
    }
    return content, nil
}

func writeSug(p string, c *string) {
    if len(strings.Split(p, "/")) > 1{
        p = "_" + p
    }
    path := filepath.Join("suggestions",p)
    dir := filepath.Dir(path)
    if err := os.MkdirAll(dir, 0755); err != nil {
        log.Printf("%v",err)
	}else{
    file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
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

	log.Println("File written successfully")
    }
}
