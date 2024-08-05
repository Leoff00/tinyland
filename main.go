package main

import (
	"flag"
	"fmt"
	"os"
)

const art string = `                                                                                                              
★ ★ ★ ★ ★ WELCOME TO TINYLAND ★ ★ ★ ★ ★
`

var (
	project string
	gpUrl   string
	root    bool
)

func main() {
	fmt.Print(art)

	flag.StringVar(&project, "project", "tinyland_example", "need to specify the name of the project.")
	flag.BoolVar(&root, "root", false, "set main.go in root dir")
	flag.StringVar(&gpUrl, "url", "github.com/tinyland_example", "need to specify your github project url")
	flag.Parse()

	attributes := FilesAttributes{
		Project: project,
		GpUrl:   gpUrl,
		Root:    root,
	}

	if gpUrl != "" {
		fmt.Printf("Setting configuration for %s \n", gpUrl)
	}

	if project != "" {
		fmt.Printf("Setting default minimalist config with %s \n", project)

		attributes.CreateProjectFolder()
		attributes.CreateGoModTemplate()
		attributes.CreateMakefileTemplate()
		attributes.CreateMainTemplate()
		clearPath(root)
	} else {
		fmt.Println("Setting default minimalist config (go.mod, Makefile and main.go)")
	}
}

func clearPath(root bool) {
	if root {
		os.Chdir("../")
		if err := os.RemoveAll("cmd"); err != nil {
			fmt.Println(err)
		}
	}

}
