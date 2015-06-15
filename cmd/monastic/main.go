package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/peterhellberg/monastic"
)

func main() {
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil {
			if c, err := monastic.Cipher(n); err == nil {
				fmt.Printf("\n%s\n", c)
			}
		}
	}
}
