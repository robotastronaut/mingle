package mpackage

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

const (
	// TODO: Detect term width and adjust
	columnWidth = 30
	lineWidth   = 80
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	list      = lipgloss.NewStyle().
			MarginRight(2).
			Height(8).
			Width(lineWidth)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Render

	listItem = lipgloss.NewStyle().PaddingLeft(2).Render
)

type MFile struct {
	Package     string `mapstructure:"package"`
	Version     string `mapstructure:"version"`
	Author      string `mapstructure:"author"`
	Description string `mapstructure:"description"`
	Title       string `mapstructure:"title"`

	_viper *viper.Viper
}

func NewMFile() MFile {
	return MFile{
		_viper: NewPackageViper("."),
	}
}

func LoadMFile() (m MFile, err error) {
	// Load the mfile using the global viper
	// TODO: Should it be the global one?
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Attempt to unmarshal the config data into our Mfile struct
	err = viper.Unmarshal(&m)
	return
}

func (m *MFile) String() string {
	return list.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			listHeader("Module Info"),
			listItem("Package:     ", m.Package),
			listItem("Version:     ", m.Version),
			listItem("Author:      ", m.Author),
			listItem("Description: ", m.Description),
			listItem("Title:       ", m.Title),
		),
	)
}

// func (m *MFile) Form() *huh.Form {
// 	form := huh.NewForm()
// }
