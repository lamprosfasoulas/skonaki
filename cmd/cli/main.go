package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/lamprosfasoulas/docs/pkg/files"
)

func main(){
    args := os.Args
    fmt.Println(string(files.GetContent(args[1:])))
    cmd := exec.Command("ls","data")
    output, _:= cmd.CombinedOutput()
    fmt.Println(strings.Fields(string(output)))
}
