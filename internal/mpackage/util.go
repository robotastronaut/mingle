package mpackage

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func FindRootFile(dir, name string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", errors.New("unable to get current user to determine home directory")
	}

	if dir == "" {
		return "", errors.New("dir must not be empty")
	}

	if name == "" {
		return "", errors.New("name must not be empty")
	}

	// Fix for home directories
	if dir == "~" {
		dir = usr.HomeDir
	}

	if strings.HasPrefix(dir, "~/") {
		dir = filepath.Join(usr.HomeDir, dir[2:])
	}

	// Parent directory
	var parent string

	// Get the absolute path to handle relative inputs
	dir, err = filepath.Abs(dir)

	if err != nil {
		return "", fmt.Errorf("unable to get absolute path while finding file: %w", err)
	}

	// Loop and look for the mfile
	for {
		// Check this dir
		f := filepath.Join(dir, name)

		// If we can stat it, we found it
		_, err := os.Stat(f)
		if err == nil {
			return dir, nil
		}

		// If not, check the parent
		parent = filepath.Dir(dir)

		// If the parent is the current dir, we've reached the root
		if parent == dir {
			return "", os.ErrNotExist
		}

		dir = parent
	}
}
