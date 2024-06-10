package muddler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Item struct {
	// Whether or not this item is active
	IsActive string `json:"isActive"`
	// Whether or not this item represents a folder
	IsFolder string `json:"isFolder"`
	// The name of the item
	Name string `json:"name"`
	// The name of the package this item belongs to
	PackageName string `json:"-"`
	// The path of the resource
	Path string `json:"-"`
	// Script holds the lua code that will be run. In the source json file, this field can be omitted to look for {item.Name}.lua
	Script string `json:"script"`
}

func (i *Item) LoadScript() error {
	// If we have something in i.Script already, bounce out
	if i.Script != "" {
		return nil
	}

	if i.Path == "" {
		return errors.New("item is missing path")
	}

	scriptPath := filepath.Join(i.Path, i.Name+".lua")
	// If we can stat it, we found it
	_, err := os.Stat(scriptPath)
	if err != nil {
		// Script file doesn't exist. This is perfectly acceptable, so don't return an error
		return nil
	}

	// Try to read the file
	data, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("unable to read script file: %w", err)
	}

	// Make sure no windows line endings appear
	i.Script = strings.Replace(string(data), "\n\r", "\n", -1)
	return nil
}

type items interface {
	Script
}

func loadItemManifestAt[T items](path string) ([]T, error) {
	file, err := os.ReadFile(filepath.Join(path, "mfile"))
	if err != nil {
		return nil, err
	}

	item := []T{}
	err = json.Unmarshal(file, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
