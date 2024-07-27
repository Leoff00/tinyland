package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var (
	fa = FilesAttributes{
		Project: "foo",
		GpUrl:   "github.com/leoff00/foo",
	}
)

func TestCreateTempFolder(t *testing.T) {
	want_cmd := "cmd"
	custom := fmt.Sprintf("%s/cmd", fa.Project)
	exist, _ := os.Stat(fa.Project)

	if exist != nil {
		os.RemoveAll(fa.Project)
	}

	fa.CreateRootFolder()
	cmd, err := os.Stat(custom)

	if cmd.Name() != want_cmd {
		t.Error("Cmd folder wasn't created", err.Error())
	}

}

func TestGoModFile(t *testing.T) {
	dir, _ := filepath.Abs(".")
	fullpath := fmt.Sprintf("%s/go.mod", dir)
	want_gomod := "go.mod"

	fa.CreateGoModTemplate()
	gomod, err := os.Stat(fullpath)

	if gomod.Name() != want_gomod {
		t.Error("go mod file wasn't created", err.Error())
	}
}

func TestCreateMainFile(t *testing.T) {
	dir, _ := filepath.Abs(".")
	fullpath := fmt.Sprintf("%s/cmd/main.go", dir)
	want_main := "main.go"

	fa.CreateMainTemplate()

	mainFile, err := os.Stat(fullpath)
	if mainFile.Name() != want_main {
		t.Error("go mod file wasn't created", err.Error())
	}

}

func TestCreateMakefile(t *testing.T) {
	dir, _ := filepath.Abs(".")
	fullpath := fmt.Sprintf("%s/Makefile", dir)
	want_makefile := "Makefile"

	fa.CreateMakefileTemplate()
	makefile, err := os.Stat(fullpath)

	if makefile.Name() != want_makefile {
		t.Error("Makefile wasn't created", err.Error())
	}
}
