package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("E:\\program\\tesseract\\tesseract.exe", "./path/ab.png", "./path/ab", "--psm", "6")

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

}