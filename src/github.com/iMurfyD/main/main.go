package main

import (
	"fmt"
	"github.com/iMurfyD/compress"
	"os"
	//"uncompress"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need a file")
	} else {
		compress.To(os.Args[1], os.Args[2])
	}
}
