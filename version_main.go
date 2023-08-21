package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	b, err := exec.Command("git", "describe", "--always", "--dirty=-modified").Output()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	s := string(b)

	s = strings.TrimSpace(s)

	t := time.Now().Format("2006-01-02 15:04:05")

	ver := "mmbee " + s + " " + t

	os.WriteFile("version.txt", []byte(ver), 0644)

	fmt.Println(ver)
}
