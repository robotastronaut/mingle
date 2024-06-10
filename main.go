/*
Copyright © 2024 Nicholas Molen <nick@robotastronaut.com>
*/

package main

import (
	"os"

	"github.com/robotastronaut/mpm/internal/cli"
)

func main() {
	err := cli.Root().Execute()
	if err != nil {
		os.Exit(1)
	}
}
