package mudlet

import "path/filepath"

// Temporary
type Package struct {
	_name string
	_path string
}

func (p Package) Name() string {
	return p._name
}

func (p Package) Path() string {
	return p._path
}

func (p Package) ConfigPath() string {
	return filepath.Join(p._path, p._name+".xml")
}

// func getPackagesInProfile(profileDir string) (packages []Package, err error) {
// 	profileDirectories, err := os.ReadDir(currentProfilePath)
// 			if err != nil {
// 				return m, fmt.Errorf("unable to get directory info for Mudlet profile %s", profileDirEntry.Name())
// 			}
// 	return
// }
