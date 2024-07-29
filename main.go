package main

import (
	"flag"
	"fmt"
)

const art string = `                                                                                                              
★ ★ ★ ★ ★ WELCOME TO TINYLAND ★ ★ ★ ★ ★
`

var (
	project string
	gpUrl   string
)

func main() {
	fmt.Print(art)

	flag.StringVar(&project, "project", "tinyland_example", "need to specify the name of the project.")
	flag.StringVar(&gpUrl, "url", "github.com/tinyland_example", "need to specify your github project url")
	flag.Parse()

	attributes := FilesAttributes{
		Project: project,
		GpUrl:   gpUrl,
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
	} else {
		fmt.Println("Setting default minimalist config (go.mod, Makefile and main.go)")
	}

}
