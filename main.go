package main

import (
	"flag"
	"os"
)

func main() {
	dirName := flag.String("dirName", "", "Choose a path to define")
	flag.Parse()

	filesTemplate := FilesTemplate{
		*dirName,
	}

	if *dirName != "" {
		os.Mkdir("tmp", 0777)
		filesTemplate.CreateGoModTemplate()
		filesTemplate.CreateMainTemplate()
		filesTemplate.CreateMakefileTemplate()
	}
}
