package muddler

import (
	"path/filepath"
)

type Script struct {
	Item
	EventHandlers []string `json:"eventHandlerList"`
	Children      []Script `json:"children"`
}

type ScriptManifest struct {
	Scripts []Script

	// The path of this file
	_path string
	// This holds child manifests located in subdirectories
	_children []ScriptManifest
}

func LoadScriptManifest(path string) (ScriptManifest, error) {
	m := ScriptManifest{
		_path: path,
	}

	scripts, err := loadItemManifestAt(path)
	if err != nil {
		return m, err
	}

	m.Scripts = scripts

	childManifestPaths, err := filepath.Glob(filepath.Dir(path) + "/*/scripts.json")
	if err != nil {
		return m, err
	}

	for _, manifestPath := range childManifestPaths {
		child, err := LoadScriptManifest(manifestPath)
		if err != nil {
			return m, err
		}

		m._children = append(m._children, child)
	}

	return m, nil
}
