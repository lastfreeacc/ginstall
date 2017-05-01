package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const maingo = "main.go"

var (
	gopath string
	gosrc  string
)

func main() {
	file := filepath.Clean(os.Args[1])
	dir, _ := filepath.Split(file)
	pre := dir[:len(gosrc)]
	if pre != gosrc {
		log.Fatalf("[Error] Current package {%s} is not in GOPATH {%s}\n", dir, gopath)
	}
	installPath := dir[len(gosrc)+1:]
	installPath = findMainDir(installPath)
	cmd := exec.Command("go", "install", installPath)
	if err := cmd.Run(); err != nil {
		log.Fatalf("[ERROR] can not run cmd 'go install %s', because of: %s\n", installPath, err)
	}
	log.Printf("[Info] Done: go install %s\n", installPath)
}

func init() {
	// required
	if len(os.Args) != 2 {
		log.Fatalf("[Error] Wrong number of parametes: %v\n", os.Args)
	}
	// set global vars
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalf("[Error] GOPATH is not set\n")
	}
	gosrc = filepath.Clean(gopath + "/src")
}

// findMainDir returns dir that contains main.go file
// if it could not found then returns input path
func findMainDir(installPath string) string {
	tryFindMain := installPath
	for len(tryFindMain) > 1 { // exit then tryFindMain = ".", TODO: need more strong condition
		fullPath := filepath.Join(gosrc, tryFindMain)
		files, err := ioutil.ReadDir(fullPath)
		if err != nil {
			log.Printf("[Info] can not find %s, because of: %s", maingo, err)
			break
		}
		for _, f := range files {
			if f.Name() == maingo {
				return tryFindMain // we find main.go
			}
		}

		tryFindMain = filepath.Join(tryFindMain, "..")
	}
	log.Printf("[Info] can not find %s\n", maingo)
	return installPath // can not find main.go
}
