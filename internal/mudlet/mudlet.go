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

			m.Profiles = append(m.Profiles, Profile{
				Name: profileDirEntry.Name(),
				Path: currentProfilePath,
			})
		}
	}

	return m, nil
}
