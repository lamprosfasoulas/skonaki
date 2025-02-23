package files

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lamprosfasoulas/docs/pkg/cache"
)

const ROOT = "data"
var DIRS []string//{"cheats", "sheets", "also"}
const prefix = "_"

func RunBat(f, lang string) ([]byte, error) {
    cmd := exec.Command("bat","--color=always","--paging=never","--style=plain",
        "-l", lang,f)
    output, err := cmd.CombinedOutput()
    if err != nil {
            return []byte{},err
    }
    return output,nil
}
func GetHTML(c []byte) []byte {
    aha := exec.Command("aha","--black","-n","-l","-t","Skonaki")
    aha.Stdin = bytes.NewReader(c)

    var ahaOut bytes.Buffer

    aha.Stdout = &ahaOut

    if err:= aha.Run(); err == nil{
        return []byte(ahaOut.String())
    }
    return []byte{}
}
func GetList() []byte {
    cmd := exec.Command("ls","-R","data")
    if dirs, err := cmd.CombinedOutput(); err == nil{
        DIRS = strings.Fields(string(dirs))
    }
    return []byte(fmt.Sprintf("%v\n",strings.Join(DIRS,"\n")))
}

func GetContent(path []string) []byte{
    var lang string //used for bat syntax highlighting
    var content string 
    var showDir string

    if path[0] == "" {
        path[0] = "home"
    }

    switch len(path){
    case 1:
        lang = "bash"
        showDir = path[0]
    case 2:
        lang = path[0]
        showDir = path[0]
        path[0] = "_" + path[0]
    }
    key := filepath.Join(path...)
    //start := time.Now()
    inner := func() {
        // Here we get the directories under data
        cmd := exec.Command("ls","data")
        if dirs, err := cmd.CombinedOutput(); err == nil{
            DIRS = strings.Fields(string(dirs))
        }
        wg := sync.WaitGroup{}
        for _,dir := range DIRS{
            wg.Add(1)
            //filePath := filepath.Join(ROOT,strings.Split(dir,".")[1],filepath.Join(path...))
            filePath := filepath.Join(ROOT,dir,filepath.Join(path...))
            //if c, e := cache.GetCont(filePath); e == nil && c != nil{
            //    fmt.Println("found in cache",filePath,e)
            //    content += string(c)
            //}else{
            if output, err := RunBat(filePath,lang); err == nil {
                tempC := fmt.Sprintf("\n\033[33;1m%v:%v\033[0m\n",strings.Split(dir,".")[1],showDir)
                tempC += string(output)
                //log.Printf("Failed to set cache: %v",cache.SetCont(filePath,[]byte(tempC)))
                content += tempC
            }
            //}
            wg.Done()
    }
    //log.Printf("\nTime for loop %v\n ----------",time.Since(start))
    if content == ""{
        GetContent([]string{"404"})
    }
    cache.SetCont(key, []byte(content))
    }
    if c, e := cache.GetCont(key); e == nil && c != nil{
        log.Printf("Getting key: %v from Cache",key)
        //go inner()
        return c
    }else{
        inner()
    }
    return []byte(content)
}


