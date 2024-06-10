package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/robotastronaut/mpm/internal/muddler"
)

type GenerateOpt uint8

func (o GenerateOpt) String() string {
	name := []string{"none", "Script", "Alias", "Trigger", "Key", "Timer"}
	i := uint8(o)
	switch {
	case i <= uint8(GenerateOptTimer):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}

const (
	_ GenerateOpt = iota
	GenerateOptScript
	GenerateOptAlias
	GenerateOptTrigger
	GenerateOptKey
	GenerateOptTimer
)

// Generator holds all the configuration needed to create a new package
type Generator struct {
	Path        string
	Name        string
	Description string
	Author      string
	Title       string
	OutputFile  bool
	Options     []GenerateOpt

	// Used to mark the generator as ready to run
	_ready bool
	// Whether or not to enter a description
	_describe bool
}

func (g *Generator) Run() error {
	rootPath, err := muddler.ResolvePath(g.Path)
	if err != nil {
		return fmt.Errorf("generator error: %w", err)
	}
	// Check if path exists
	_, err = os.Stat(rootPath)
	if os.IsNotExist(err) {
		// create the module path
		err = os.MkdirAll(rootPath, 0777)
		if err != nil {
			return fmt.Errorf("generator failed to create package directory: %w", err)
		}
	}

	if g.Name == "" {
		g.Name = "changeme"
	}
	if g.Title == "" {
		g.Title = "changeme"
	}

	// Generate mfile
	mfile := muddler.Module{
		Package:     g.Name,
		Version:     "0.0.1",
		Author:      g.Author,
		Title:       g.Title,
		OutputFile:  g.OutputFile,
		Path:        rootPath,
		Description: g.Description,
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
		), huh.NewGroup(
			huh.NewInput().
				Title("Title").
				Description("eg 'Personal Mudlet Package'").
				Value(&g.Title),
		), huh.NewGroup(
			huh.NewInput().
				Title("Author").
				Description("eg 'Nick Molen <nick@robotastronaut.com>'").
				Value(&g.Author),
		), huh.NewGroup(
			huh.NewConfirm().
				Title("Would you like to enter a description for this package?").
				Value(&g._describe),
		), huh.NewGroup(
			huh.NewText().
				Title("Description").
				Value(&g.Description),
		).WithHideFunc(func() bool { return !g._describe }), huh.NewGroup(
			huh.NewConfirm().
				Title("Muddler File Watching").
				Description("If enabled, your package will generate a .output file when built.").
				Value(&g.OutputFile),
		), huh.NewGroup(
			huh.NewMultiSelect[GenerateOpt]().
				Options(
					huh.NewOption("Scripts", GenerateOptScript),
					huh.NewOption("Aliases", GenerateOptAlias),
					huh.NewOption("Triggers", GenerateOptTrigger),
					huh.NewOption("Keys", GenerateOptKey),
					huh.NewOption("Timers", GenerateOptTimer),
					// huh.NewOption("Enable File Watching", GenerateOptOutputFile),
				).
				Title("Example Generation").
				Description("Select examples to be generated").
				Value(&g.Options),
		),
	).WithTheme(theme)
}

func (g *Generator) Confirm() *huh.Form {
	g._ready = true
	fmt.Println(g.Summary())
	return huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Generate package with these settings?").
				Affirmative("Generate").
				Negative("Cancel").
				Value(&g._ready),
		),
	).WithTheme(theme)
}

func (g *Generator) Summary() string {
	summary := strings.Builder{}

	summary.WriteString(summaryLine("Path", g.Path))
	summary.WriteString(summaryLine("Name", g.Name))
	summary.WriteString(summaryLine("Title", g.Title))
	summary.WriteString(summaryLine("Author", g.Author))
	summary.WriteString(summaryLine("Use .output", strconv.FormatBool(g.OutputFile)))
	optValues := []string{}
	for _, opt := range g.Options {
		optValues = append(optValues, opt.String())
	}

	summary.WriteString(summaryLine("Examples", strings.Join(optValues, ", ")))
	summary.WriteString(summaryBlock("Description", g.Description))

	return container.Render(summary.String())
}
