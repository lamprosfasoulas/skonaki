package files

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

const ROOT = "data"
var DIRS = []string{"cheats", "sheets", "also"}
const prefix = "_"

func runBat(f, lang string) ([]byte, error) {
    cmd := exec.Command("bat","--color=always","--paging=never","--style=plain",
        "-l", lang,f)
    output, err := cmd.CombinedOutput()
    if err != nil {
            return []byte{},err
    }
    return output,nil
}
func GetHTML(c []byte) []byte {
    aha := exec.Command("aha","--black")
    aha.Stdin = bytes.NewReader(c)
    var out bytes.Buffer

    aha.Stdout = &out

    if err:= aha.Run(); err == nil{
        return []byte(out.String())
    }
    return []byte{}

}

func GetContent(path []string) []byte{
    var lang string //used for bat syntax highlighting
    var content string 
    // cmd := exec.Command("bash","-c","gls","-d","data/*")
    // if out, err := cmd.CombinedOutput(); err == nil{
    //     DIRS = strings.Fields(string(out))
    // }
    fmt.Println(DIRS)
    fmt.Println(path)
    switch len(path){
    case 1:
        lang = "bash"
    case 2:
        lang = path[0]
        path[0] = "_" + path[0]
    }
    for _,dir := range DIRS{
        filePath := filepath.Join(ROOT,dir,filepath.Join(path...))
        if output, err := runBat(filePath,lang); err == nil{
            fmt.Println(filePath)
            content += (fmt.Sprintf("\n\033[33;1m%v:%v\033[0m\n",dir,path[0]))
            content += string(output)
        }
    }
    if content == ""{
        content = "No command provided\n"
    }
    return []byte(content)
}


