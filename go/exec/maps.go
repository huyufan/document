package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	goBin, _ := filepath.Abs(filepath.Join("build", "bin", "cc"))

	fmt.Println(goBin)

	dir, _ := os.Getwd()
	fmt.Println(dir)

}
