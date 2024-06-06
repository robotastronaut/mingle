package cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/robotastronaut/mingle/internal/muddler"
)

type GenerateOpt uint8

const (
	_ GenerateOpt = iota
	GenerateOptScript
	GenerateOptAlias
	GenerateOptTrigger
	GenerateOptKey
	GenerateOptTimer
	GenerateOptOutputFile
)

// Generator holds all the configuration needed to create a new package
type Generator struct {
	Path    string
	Name    string
	Author  string
	Title   string
	Options []GenerateOpt
}

func (g *Generator) Run() error {
	rootPath, err := muddler.ResolvePath(g.Path)
	fmt.Println("GENERATING AT", rootPath)
	if err != nil {
		return fmt.Errorf("generator error: %w", err)
	}
	fmt.Println("GENERATING AT", rootPath)
	// Check if path exists
	_, err = os.Stat(rootPath)
	if os.IsNotExist(err) {
		// create the module path
		err = os.MkdirAll(rootPath, 0777)
		if err != nil {
			return fmt.Errorf("generator failed to create package directory: %w", err)
		}
	}

	// Generate mfile
	mfile := muddler.Module{
		Package:    g.Name,
		Version:    "0.0.1",
		Author:     g.Author,
		Title:      g.Title,
		OutputFile: true,
		Path:       rootPath,
	}

	err = mfile.Save()
	if err != nil {
		return fmt.Errorf("generator failed to write mfile: %w", err)
	}

	return nil
}

func (g *Generator) Form() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Description("Should only contain letters, numbers, underscores, and dashes").
				Validate(nameValidator).
				Value(&g.Name),
			huh.NewInput().
				Title("Title").
				Description("eg 'Personal Mudlet Package'").
				Value(&g.Title),
			huh.NewInput().
				Title("Author").
				Description("eg 'Nick Molen <nick@robotastronaut.com>'").
				Value(&g.Author),
		), huh.NewGroup(
			huh.NewMultiSelect[GenerateOpt]().
				Options(
					huh.NewOption("Example Scripts", GenerateOptScript),
					huh.NewOption("Example Aliases", GenerateOptAlias),
					huh.NewOption("Example Triggers", GenerateOptTrigger),
					huh.NewOption("Example Keys", GenerateOptKey),
					huh.NewOption("Example Timers", GenerateOptTimer),
					huh.NewOption("Enable File Watching", GenerateOptOutputFile),
				).
				Title("Generation Options").
				Value(&g.Options),
		), huh.NewGroup(
			huh.NewNote().Title("TEST NOTE").Description("TEST DESCRIPTION"),
		),
	)
}
