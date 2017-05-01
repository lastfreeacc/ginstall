package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const maingo = "main.go"

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("[Error] Wrong number of parametes: %v\n", os.Args)
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalf("[Error] GOPATH is not set\n")
	}
	gosrc := filepath.Clean(gopath + "/src")
	file := filepath.Clean(os.Args[1])
	dir, _ := filepath.Split(file)
	pre := dir[:len(gosrc)]
	if pre != gosrc {
		log.Fatalf("[Error] Current package {%s} is not in GOPATH {%s}\n", dir, gopath)
	}
	installPath := dir[len(gosrc)+1:]
	tryFindMain := installPath
OUTER:
	for len(tryFindMain) > 1 {
		fullPath := filepath.Join(gosrc, tryFindMain)
		files, err := ioutil.ReadDir(fullPath)
		if err != nil {
			log.Printf("[Info] can not find %s, because of %s", maingo, err)
			break
		}
		for _, f := range files {
			if f.Name() == maingo {
				installPath = tryFindMain
				break OUTER
			}
		}

		tryFindMain = filepath.Join(tryFindMain, "..")
	}
	cmd := exec.Command("go", "install", installPath)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Printf("[Info] Done: go install %s\n", installPath)
}
