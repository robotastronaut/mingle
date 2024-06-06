package muddler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Module struct {
	Package     string `json:"package"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Title       string `json:"title"`
	OutputFile  bool   `json:"outputFile"`

	// A filesystem interface that should be rooted at `_path`
	Path string `json:"-"`
}

func (m *Module) Save() error {
	d, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(m.Path, "mfile"), d, 0666)
	if err != nil {
		return err
	}

	return nil
}

func loadModuleAt(path string) (*Module, error) {

	file, err := os.ReadFile(filepath.Join(path, "mfile"))
	if err != nil {
		return nil, err
	}

	m := Module{}
	err = json.Unmarshal(file, &m)
	if err != nil {
		return nil, err
	}

	m.Path = path
	return &m, nil
}

func FindModule(dir string) (*Module, error) {
	modPath, err := FindRootFile(dir, "mfile")
	if err != nil {
		return nil, fmt.Errorf("failed to locate module: %w", err)
	}

	// Go ahead and load the module
	m, err := loadModuleAt(modPath)

	if err != nil {
		return m, fmt.Errorf("error loading module: %w", err)
	}

	return m, nil

}
