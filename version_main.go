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

	// 获取golang版本
	v := runtime.Version()

	// 将golang版本添加到ver变量中
	ver := "mmbee " + s + " " + t + " " + v

	os.WriteFile("version.txt", []byte(ver), 0644)

	fmt.Println(ver)
}
