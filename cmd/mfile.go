package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"golang.org/x/mod/semver"
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

	packageRegex = regexp.MustCompile("^[a-zA-Z0-9]+$")
)

type InitExampleOpt uint8

const (
	_ InitExampleOpt = iota
	InitOptScript
	InitOptAlias
	InitOptTrigger
	InitOptKey
	InitOptTimer
)

type Mfile struct {
	Package     string `mapstructure:"package"`
	Version     string `mapstructure:"version"`
	Author      string `mapstructure:"author"`
	Description string `mapstructure:"description"`
	Title       string `mapstructure:"title"`
	OutputFile  bool   `mapstructure:"outputFile"`

	InitOpts []InitExampleOpt
	_viper   *viper.Viper
}

func (m *Mfile) Save() error {
	if m == nil {
		return fmt.Errorf("error saving mfile: nil mfile")
	}

	if m._viper == nil {
		return fmt.Errorf("error saving mfile: nil viper")
	}

	m._viper.WriteConfig()
	return nil
}

func (m *Mfile) String() string {
	return list.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			listHeader("Module Info"),
			listItem("Package:     ", m.Package),
			listItem("Version:     ", m.Version),
			listItem("Author:      ", m.Author),
			listItem("Description: ", m.Description),
			listItem("Title:       ", m.Title),
			listItem("OutputFile:  ", strconv.FormatBool(m.OutputFile)),
		),
	)
}

// Form

func (m *Mfile) Form() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Description("Should only contain letters, numbers, underscores, and dashes").
				// Validate(func(t string) error {
				// 	if !packageRegex.MatchString(t) {
				// 		return fmt.Errorf("invalid package name")
				// 	}
				// 	return nil
				// }).
				Value(&m.Package),
			huh.NewInput().
				Title("Title").
				Description("eg 'Personal Mudlet Package'").
				Value(&m.Title),
			huh.NewInput().
				Title("Author").
				Description("eg 'Nick Molen <nick@robotastronaut.com>'").
				Value(&m.Author),
			huh.NewInput().
				Title("Version").
				Description("Must be valid SemVer (eg '0.0.1')").
				Validate(func(v string) error {
					if !semver.IsValid("v" + v) {
						return fmt.Errorf("invalid semantic version")
					}
					return nil
				}).
				Value(&m.Version),
			huh.NewText().
				Title("Description").
				Value(&m.Description),
		),
		huh.NewGroup(
			huh.NewMultiSelect[InitExampleOpt]().
				Title("Select examples to create in your package").
				Options(
					huh.NewOption("Script", InitOptScript),
					huh.NewOption("Alias", InitOptAlias),
					huh.NewOption("Trigger", InitOptTrigger),
					huh.NewOption("Key", InitOptKey),
					huh.NewOption("Timer", InitOptTimer),
				).
				Value(&m.InitOpts),
		),
	)
}

func NewMfile(path string) Mfile {
	v := newMfileViper()
	v.AddConfigPath(path)
	return Mfile{
		_viper: v,
	}
}

func LoadMfile(path string) (m Mfile, err error) {
	m._viper = newMfileViper()
	m._viper.AddConfigPath(path)
	// Load the mfile
	err = m._viper.ReadInConfig()
	if err != nil {
		return
	}

	// Attempt to unmarshal the config data into our Mfile struct
	err = m._viper.Unmarshal(&m)
	return
}

func newMfileViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("mfile")
	v.SetConfigType("json")

	return v
}

func getModulePath(args []string) (string, error) {
	modPath := "."
	if len(args) > 0 && len(args[0]) > 0 {
		modPath = args[0]
	}

	_, err := os.Stat(modPath)
	if err == nil {
		return modPath, nil
	}

	return modPath, err
}
