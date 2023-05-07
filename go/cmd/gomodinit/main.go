package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	_, file, _, _ := runtime.Caller(0)
	root := filepath.Dir(filepath.Dir(filepath.Dir(file)))
	for _, year := range readDir(root) {
		if year.IsDir() && len(year.Name()) == 4 {
			goModInit(filepath.Join(root, year.Name()))
			for _, day := range readDir(filepath.Join(root, year.Name())) {
				if strings.HasPrefix(day.Name(), "day") {
					goModInit(filepath.Join(root, year.Name(), day.Name()))
				}
			}
		}
	}
}
func readDir(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Panicln(err)
	}
	return entries
}
func goModInit(dir string) {
	log.Println("running `go mod init && go mod tidy && go work use .` at:", dir)
	runAt(dir, "go", "mod", "init")
	runAt(dir, "go", "mod", "tidy")
	runAt(dir, "go", "work", "use", ".")
}
func runAt(dir, command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Panicln(err)
	}
}
