package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	createPot()
}


func createPot() {
        if !Exists("/usr/local/bin/gettext") {
                fmt.Println("gettext does not exist")
        }
        if Exists("/usr/local/bin/gettext") {
                execmd := exec.Command("xgettext","--from-code=UTF-8", "--files-from=./files_to_translate", "--keyword=translatable", "--keyword=_", "--output=../../po/ghostbsd-welcome.pot")
		execmd.Start()
        }
}


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		checkError(err)
	}
	return true
}
