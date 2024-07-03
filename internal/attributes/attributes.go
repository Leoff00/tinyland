package attributes

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type FilesAttributes struct {
	//The whole project including go.mod, dir and self

	Project string
}

func (f *FilesAttributes) truncateVersion() string {
	numbers := strings.Split(runtime.Version(), "go")
	return strings.TrimLeft(strings.Join(numbers, ""), "")

}

func (f *FilesAttributes) CreateMakefileTemplate() {

	t := fmt.Sprintln(`
run:
		@go run cmd/main.go`)

	file, err := os.Create("tmp/Makefile")

	if err != nil {
		log.Panicln("Erro during create file Makefile", err.Error())
	}

	defer file.Close()

	if _, err := file.WriteString(t); err != nil {
		log.Println("Could not write the content in file...", err.Error())
	}

}

func (f *FilesAttributes) CreateGoModTemplate() {
	t := fmt.Sprintf(`module %s

go %s`, f.Project, f.truncateVersion(),
	)

	file, err := os.Create("tmp/go.mod")

	if err != nil {
		log.Panicln("Erro during create file go.mod", err.Error())
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

	file, err := os.Create("tmp/cmd/main.go")

	if err != nil {
		log.Panicln("Erro during create file main.go", err.Error())
	}

	defer file.Close()

	if _, err := file.WriteString(t); err != nil {
		log.Println("Could not write the content in file...", err.Error())
	}
}