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
		Root:    true,
	}
)

func TestCreateProjectFolder(t *testing.T) {
	want_cmd := "cmd"
	custom := fmt.Sprintf("%s/cmd", fa.Project)
	exist, _ := os.Stat(fa.Project)

	if exist != nil {
		os.RemoveAll(fa.Project)
	}

	fa.CreateProjectFolder()
	cmd, err := os.Stat(custom)

	if cmd.Name() != want_cmd {
		t.Error("Cmd folder wasn't created", err.Error())
	}

}

// abs used because in runtime is entering in folder created...

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

func TestCreateMainFile(t *testing.T) {
	if fa.Root {
		dir, _ := filepath.Abs(".")
		fullpath := fmt.Sprintf("%s/main.go", dir)
		want_main := "main.go"

		fa.CreateMainTemplate()

		mainFile, err := os.Stat(fullpath)
		if mainFile.Name() != want_main {
			t.Error("main file wasn't created", err.Error())
		}
	} else {
		dir, _ := filepath.Abs(".")
		fullpath := fmt.Sprintf("%s/cmd/main.go", dir)
		want_main := "main.go"

		fa.CreateMainTemplate()

		mainFile, err := os.Stat(fullpath)
		if mainFile.Name() != want_main {
			t.Error("main file wasn't created", err.Error())
		}
	}
}
