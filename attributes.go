package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type FilesAttributes struct {
	//The whole project including go.mod, dir and self
	Project string

	// the github profile url of the user
	GpUrl string
}

func (f *FilesAttributes) CreateProjectFolder() {
	path := fmt.Sprintf("%s/cmd", f.Project)
	if err := os.MkdirAll(path, 0777); err != nil {
		fmt.Printf("fail to create folder %s, error: %s\n", f.Project, err.Error())
	}

	if err := os.Chmod(f.Project, 0777); err != nil {
		fmt.Printf("fail to give perms to %s, error: %s\n", f.Project, err.Error())
	}
}

func (f *FilesAttributes) CreateGoModTemplate() {
	if err := os.Chdir(f.Project); err != nil {
		fmt.Printf("Fail to enter %s directory, error: %s\n", f.Project, err.Error())
	}

	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		gomodCmd := exec.Command("go", "mod", "init", f.GpUrl)
		if err := gomodCmd.Run(); err != nil {
			if err = os.RemoveAll(f.Project); err != nil {
				fmt.Printf("Error during removing %s directory, error: %s\n", f.Project, err.Error())
			}
			log.Panicln("Fail to execute go mod command...", err.Error())
		}
	} else {
		return
	}
}

func (f *FilesAttributes) CreateMakefileTemplate() {

	t := fmt.Sprintln(`
run:
		@go run cmd/main.go`)

	file, err := os.Create("Makefile")

	if err != nil {
		log.Panicln("Erro during create file Makefile", err.Error())
	}

	defer file.Close()

	if _, err := file.WriteString(t); err != nil {
		log.Println("Could not write the content in file...", err.Error())
	}
}

func (f *FilesAttributes) CreateMainTemplate() {
	t := fmt.Sprintln(`package main

import "fmt"

func main() {
	fmt.Println("hello world")
}`)

	if err := os.Chdir("cmd"); err != nil {
		os.RemoveAll(f.Project)
		log.Panicln("Fail in enter cmd directory...", err.Error())
	}

	file, err := os.Create("main.go")

	if err != nil {
		log.Panicln("Erro during create file main.go", err.Error())
	}

	defer file.Close()

	if _, err := file.WriteString(t); err != nil {
		log.Println("Could not write the content in file...", err.Error())
	}
}
