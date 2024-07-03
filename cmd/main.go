package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leoff00/tinyland/internal/attributes"
)

const art string = `                                                                                                              
★ ★ ★ ★ ★ WELCOME TO TINY ISLAND ★ ★ ★ ★ ★
`

var (
	project string
)

func main() {
	fmt.Print(art)

	flag.StringVar(&project, "project", "tinyland_example", "need to specify the name of the project.")
	flag.Parse()

	attributes := attributes.FilesAttributes{
		Project: project,
	}

	if project != "" {
		fmt.Printf("Setting default minimalist config with %s \n", project)
		os.Mkdir("tmp/cmd", 0777)

		attributes.CreateGoModTemplate()
		attributes.CreateMainTemplate()
		attributes.CreateMakefileTemplate()
	} else {
		fmt.Println("Setting default minimalist config (go.mod, Makefile and main.go)")
	}

}
