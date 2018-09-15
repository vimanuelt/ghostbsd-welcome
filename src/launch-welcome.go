// This is the ghostbsd-welcome launch program.
// It replaces the original launch.sh script.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	startWelcome()
}


func startWelcome() {
        if !Exists("/usr/local/bin/python3.6") {
                fmt.Println("python3.6 does not exist")
        }
        if Exists("/usr/local/bin/python3.6") {
                execmd := exec.Command("/usr/local/bin/python3.6", "render/ghostbsd-welcome.py", "--dev")
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
