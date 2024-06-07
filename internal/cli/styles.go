package cli

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/charmbracelet/bubbles/runeutil"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/rogpeppe/go-internal/semver"
)

var (
	packageRegex = regexp.MustCompile("^[a-zA-Z0-9_-]+$")
)

var (
	theme     = huh.ThemeBase16()
	container = lipgloss.
			NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("9")).
			Padding(1)

	header = theme.Focused.Title.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Color("6"))
)

func summaryLine(key, value string) string {
	return fmt.Sprintf("%-20s %s\n", theme.Focused.Title.Render(key), value)
}

func summaryBlock(key, value string) string {
	return fmt.Sprintf("%s\n%s\n", theme.Focused.Title.Render(key), value)
}

func nameValidator(s string) error {
	if !packageRegex.MatchString(s) {
		return errors.New("invalid name")
	}

	return nil
}

func semverValidator(s string) error {
	if !semver.IsValid(s) {
		return errors.New("invalid semver")
	}

	return nil
}

func blankSanitizer() runeutil.Sanitizer {
	return runeutil.NewSanitizer(
		runeutil.ReplaceTabs(""), runeutil.ReplaceNewlines(""))
}
