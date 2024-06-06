package cli

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/charmbracelet/bubbles/runeutil"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/mod/semver"
)

type (
	errMsg error
)

const (
	pkgInput = iota
	versionInput
	authorInput
	descInput
	titleInput
)

var (
	packageRegex = regexp.MustCompile("^[a-zA-Z0-9_-]+$")
)

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

type tuiModel struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func initialTUIModel() tuiModel {
	inputs := make([]textinput.Model, 3)
	inputs[pkgInput] = textinput.New()
	inputs[pkgInput].Placeholder = "package-name"
	inputs[pkgInput].Focus()
	inputs[pkgInput].CharLimit = 30
	inputs[pkgInput].Width = 40
	inputs[pkgInput].Prompt = ""
	inputs[pkgInput].Validate = nameValidator
	inputs[pkgInput].Sanitizer = blankSanitizer()

	inputs[versionInput] = textinput.New()
	inputs[versionInput].Placeholder = "0.0.1"
	inputs[versionInput].CharLimit = 10
	inputs[versionInput].Width = 40
	inputs[versionInput].Prompt = ""
	inputs[versionInput].Validate = semverValidator
	inputs[versionInput].Sanitizer = blankSanitizer()

	inputs[authorInput] = textinput.New()
	inputs[authorInput].Placeholder = "Your Name <you@yourdomain.com>"
	inputs[authorInput].CharLimit = 40
	inputs[authorInput].Width = 40
	inputs[authorInput].Prompt = ""
	inputs[authorInput].Sanitizer = blankSanitizer()

	return tuiModel{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

// Init
func (m tuiModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update
func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// View

func (m tuiModel) View() string {
	return fmt.Sprintf(
		` Total: $21.50:

 %s  %s
 %s  %s
 %s  %s

 %s
`,
		inputStyle.Width(10).Render("Name"),
		m.inputs[pkgInput].View(),
		inputStyle.Width(10).Render("Version"),
		m.inputs[versionInput].View(),
		inputStyle.Width(10).Render("Author"),
		m.inputs[authorInput].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

// Helpers

func (m *tuiModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *tuiModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
