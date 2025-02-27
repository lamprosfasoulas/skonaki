package files

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lamprosfasoulas/skonaki/pkg/cache"
)

const ROOT = "data"
var DIRS []string//{"cheats", "sheets", "also"}
const prefix = "_"
var lange = []string{
    "python", "javascript", "c", "c++", "go", "rust", "java", "ruby", "php", 
    "bash", "zsh", "perl", "lua", "swift", "typescript", "kotlin", "html", 
    "css", "shell scripts", "yaml", "json", "toml", "markdown", "latex", 
    "sql", "haskell", "r", "dart", "c#", "xml", "csv", "markdown", 
    "toml", "text", "ini", "gitignore", "dockerfile", "docker", "vim", "nginx", 
    "cmake", "latex", "shell", "racket", "ocaml", "tcl", "scheme", 
    "hcl", "yaml", "makefile", "vhdl", "vala", "graphql", "actionscript", 
    "turing", "julia", "puppet", "postscript", "sass", "scala", "haxe", 
    "swift", "css", "xslt", "qml", "viml", "vhdl", "toml", "d", 
    "nim", "fortran", "gdscript", "dart", "f#", "awk", "vbscript", 
    "idl", "crystal", "autohotkey", "openscad", "rust", "groovy", "ocaml",
}


//func RunBat(f, lang string) ([]byte, error) {
//    cmd := exec.Command("bat","--color=always","--paging=never","--style=plain",
//        "-l", lang,f)
//    output, err := cmd.CombinedOutput()
//    if err != nil {
//        return []byte{},err
//    }
//    return output,nil
//}
func RunBat(f *[]byte, lang string) []byte {
    cmd := exec.Command("bat","--color=always","--paging=never","--style=plain",
        "-l", lang)
    cmd.Stdin = bytes.NewReader(*f)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return []byte{}
    }
    return output
}
func GetHTML(c *[]byte) []byte {
    aha := exec.Command("aha","--black","-n","-l","-t","Skonaki")
    aha.Stdin = bytes.NewReader(*c)

    //var ahaOut bytes.Buffer

    //aha.Stdout = &ahaOut

    if ahaOut, err:= aha.CombinedOutput(); err == nil{
        return []byte(ahaOut)
    }
    return []byte{}
}

func GetContent(path []string) *[]byte{
    key := filepath.Join(path...)
    if c, e := cache.GetCont(key); e == nil && c != nil{
        log.Printf("Getting key: %v from Cache",key)
        return &c
    }else{
        log.Printf("Searching %v ...",path)
        resp := inner(path)
        cache.SetCont(key, resp)
        return &resp
    }
}
type Item struct {
    id  int
    p   string
}

func inner(path []string) []byte {
    var lang string //used for bat syntax highlighting
    var content string 
    var showDir string

    if path[0] == "" {
        path[0] = "home"
    }
    if path[0] == ":list"{
        return list()
    }

    switch len(path){
    case 1:
        lang = "bash"
        showDir = path[0]
    case 2:
        for _, v := range lange{
            if v == path[0]{
                lang = path[0]
                break
            }else{
                lang = "bash"
            }
        }
        showDir = path[0]
        path[0] = "_" + path[0]
    }
    var wg sync.WaitGroup
    ch := make(chan Item,len(DIRS))
    for i,dir := range DIRS{
        wg.Add(1)
        go func(id int, dir string){
            defer wg.Done()
            var res Item
            res.id = id
            filePath := filepath.Join(ROOT,dir,filepath.Join(path...))
            if output, err := os.ReadFile(filePath);err == nil {
                tempC := fmt.Sprintf("\n\033[33;1m%v:%v\033[0m\n",strings.Split(dir,".")[1],showDir)
                tempC += string(RunBat(&output,lang))
                res.p =  tempC + "\n"
                ch <- res
            }        
        }(i,dir)
    }
    go func(){
        wg.Wait()
        close(ch)
    }()
    // problem for tommorow (( order of the go routine output ))
//   for _,dir := range DIRS{
//       filePath := filepath.Join(ROOT,dir,filepath.Join(path...))
//       if output, err := RunBat(filePath,lang);err == nil {
//           tempC := fmt.Sprintf("\n\033[33;1m%v:%v\033[0m\n",strings.Split(dir,".")[1],showDir)
//           tempC += string(output)
//           content += tempC + "\n"
//       }        
//   }
    page := make([]string, len(DIRS))
    for v := range ch{ //concat based on the index order
        page[v.id] = v.p
    }
    for _, v := range page {
        content += v
    }
    if content == "" {
        wg.Wait()
        return inner([]string{"404"})
    }
    return []byte(content)
}

func list() []byte{
    var contList string
    err := filepath.Walk(ROOT, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            log.Println(err)
            return nil
        }
        if !info.IsDir(){
            contList += fmt.Sprintf("%v\n",strings.Split(path,".")[1])
        }
        return nil

    })
    if err != nil{ return nil }
     return []byte(contList)
}
