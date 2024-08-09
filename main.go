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
	url     string
	help    bool
	root    bool
)

func main() {

	flag.StringVar(&project, "project", "", "need to specify the name of the project.")
	flag.StringVar(&url, "url", "", "need to specify your github project url.")
	flag.BoolVar(&root, "root", false, "if active this flag sets the main.go to be created in root dir.")
	flag.BoolVar(&help, "h", false, "help flag to get infos")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("Examples:")
		fmt.Println("  - To create a project with a specified name:")
		fmt.Println("    $ myprogram -project=myproject -url=https://github.com/user/repo")
		fmt.Println("  - To set the main.go in the root directory:")
		fmt.Println("    $ myprogram -project=myproject -url=https://github.com/user/repo -root=true")
	}
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if project == "" && url == "" {
		flag.Usage()
		os.Exit(1)
	}

	if project != "" && url != "" {
		attributes := FilesAttributes{
			Project: project,
			Url:     url,
			Root:    root,
		}

		fmt.Print(art)
		fmt.Printf("Setting default minimalist config with %s \n", project)

		attributes.CreateProjectFolder()
		attributes.CreateGoModTemplate()
		attributes.CreateMakefileTemplate()
		attributes.CreateMainTemplate()
		clearPath(root)
	}
}

func clearPath(root bool) {
	if root {
		if err := os.RemoveAll("cmd"); err != nil {
			fmt.Println(err)
		}
	}

}
