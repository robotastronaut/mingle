package cli

import "github.com/charmbracelet/lipgloss"

const (
	// TODO: Detect term width and adjust
	columnWidth = 30
	lineWidth   = 80
)

var (
	// Colors
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	hotPink   = lipgloss.Color("#FF06B7")
	darkGray  = lipgloss.Color("#767676")

	// Basic styles
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)

	list = lipgloss.NewStyle().
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
