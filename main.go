/*
Copyright © 2024 Nicholas Molen <nick@robotastronaut.com>
*/

package main

import (
	"os"

	"github.com/robotastronaut/muddler-go/cmd"
)

func main() {
	err := cmd.Root().Execute()
	if err != nil {
		os.Exit(1)
	}
}
