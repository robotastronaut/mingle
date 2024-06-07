package mudlet

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

type Profile struct {
	Name     string
	Path     string
	Packages []Package
}

type Mudlet struct {
	ConfigPath string
	Profiles   []Profile
}

func GetMudlet() (Mudlet, error) {
	m := Mudlet{}
	usr, err := user.Current()
	if err != nil {
		return m, errors.New("unable to get current user to determine home directory")
	}

	mudletPath := filepath.Join(usr.HomeDir, ".config/mudlet")
	if _, err = os.Stat(mudletPath); err != nil {
		return m, errors.New("unable to locate Mudlet configuration directory")
	}

	m.ConfigPath = mudletPath
	profilesPath := filepath.Join(mudletPath, "profiles")
	// Get profiles
	profilesDirContents, err := os.ReadDir(profilesPath)
	if err != nil {
		return m, errors.New("unable to access Mudlet profiles")
	}

	for _, profileDirEntry := range profilesDirContents {
		if profileDirEntry.IsDir() {
			// Get Packages
			currentProfilePath := filepath.Join(profilesPath, profileDirEntry.Name())
			prof := Profile{
				Name: profileDirEntry.Name(),
				Path: currentProfilePath,
			}

			// Find packages
			profileContents, err := os.ReadDir(currentProfilePath)
			if err != nil {
				return m, errors.New("unable to access Mudlet profile contents")
			}

			for _, profileItem := range profileContents {
				if profileItem.IsDir() {
					// check if it has an xml file matching the file path
					packagePath := filepath.Join(currentProfilePath, profileItem.Name())
					xmlPath := filepath.Join(packagePath, profileItem.Name()+".xml")
					if _, err := os.Stat(xmlPath); err == nil {
						prof.Packages = append(prof.Packages, Package{
							_name: profileItem.Name(),
							_path: packagePath,
						})
					}
				}
			}

			m.Profiles = append(m.Profiles, prof)
		}
	}

	return m, nil
}
