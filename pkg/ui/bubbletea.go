package ui

import (
	"errors"
	"strings"

	"github.com/sansmoraxz/toi-go/pkg/game"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type UI struct {
	game *game.Hanoi
	currentPeg rune

	err error
	done bool
	quitting bool
	rules bool

	help help.Model

	screenWidth int
	screenHeight int
}

var (
	currentPegStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	normalPegStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	rulesStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("220"))
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)

// keyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	Quit  key.Binding
	Help  key.Binding
	Rules key.Binding
	Left  key.Binding
	Right key.Binding
	R  key.Binding
	A  key.Binding
	D  key.Binding
}

func (k keyMap) Map() map[string]key.Binding {
	return map[string]key.Binding{
		"q": k.Quit,
		"?": k.Help,
		"r": k.R,
		"f": k.Rules,
		"a": k.A,
		"d": k.D,
		"←": k.Left,
		"→": k.Right,
	}
}

var keys = keyMap{
	Quit:  key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "Quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Expand help"),
	),
	Rules: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "View rules"),
	),
	Left:  key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "Move cursor left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "Move cursor right"),
	),
	R: key.NewBinding(
		key.WithKeys("r", "R"),
		key.WithHelp("r", "Reload"),
	),
	A: key.NewBinding(
		key.WithKeys("a", "A"),
		key.WithHelp("a", "Move disk left"),
	),
	D: key.NewBinding(
		key.WithKeys("d", "D"),
		key.WithHelp("d", "Move disk right"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.R, k.Quit, k.Help, k.Rules}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.R, k.Quit, k.Help, k.Rules}, // first column
		{k.Left, k.Right},				// first column
		{k.A, k.D}, 					// third column
	}
}


func NewUI() *UI {
	hanoi := game.NewHanoi()
	return &UI{
		game: hanoi,
		currentPeg: 0,

		err: nil,
		done: false,
		quitting: false,
		rules: false,

		help: help.New(),

		screenWidth: 80,
		screenHeight: 80,
	}
}

func (ui *UI) Start() error {
	p := tea.NewProgram(
		ui,
		tea.WithAltScreen(),
	)
	_, err := p.Run()
	return err
}

func (ui *UI) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var flags tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ui.screenWidth = msg.Width
		ui.screenHeight = msg.Height
		// If we set a width on the help menu it can gracefully truncate its view as needed.
		ui.help.Width = msg.Width
		flags = tea.ClearScreen
	case tea.KeyMsg:
		flags = tea.ClearScreen
		switch {
		case key.Matches(msg, keys.Quit):
			ui.quitting = true
			return ui, tea.Quit
		case key.Matches(msg, keys.Help):
			ui.help.ShowAll = !ui.help.ShowAll
		case key.Matches(msg, keys.R):
			ui.game.Reset()
			ui.err = errors.New("game reset")
		case key.Matches(msg, keys.Rules):
			ui.rules = !ui.rules
		// arrow keys to move the cursor
		case key.Matches(msg, keys.Left):
			ui.currentPeg = game.PrevPeg(ui.currentPeg)
		case key.Matches(msg, keys.Right):
			ui.currentPeg = game.NextPeg(ui.currentPeg)

		// a or d to move disk
		case key.Matches(msg, keys.A):
			newPeg := game.PrevPeg(ui.currentPeg)
			ui.err = ui.game.MoveDisk(ui.currentPeg, newPeg)
			if ui.err != nil {
				ui.currentPeg = newPeg
			}
		case key.Matches(msg, keys.D):
			newPeg := game.NextPeg(ui.currentPeg)
			ui.err = ui.game.MoveDisk(ui.currentPeg, newPeg)
			if ui.err != nil {
				ui.currentPeg = newPeg
			}
		}
	}
	ui.done = ui.game.IsFinished()

	return ui, flags
}

func ViewRules() string {
	s := ""
	rulesStyle.Bold(true).Underline(true)
	s += rulesStyle.Render("Rules:")
	rulesStyle.Bold(false).Underline(false)
	s += rulesStyle.Render(
		"\n\n" +
		"1. Only one disk can be moved at a time.\n" +
		"2. Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack.\n" +
		"3. No disk may be placed on top of a smaller disk.\n" +
		"4. Fill the last peg with all the disks to win.\n" +
		"5. You can only move adjacent disks.\n")
	return s
}

func (ui *UI) viewBoard() string {
	var style lipgloss.Style
	s := ""
	// vertically stacked
	for i := game.NDsks - 1; i >= 0; i-- {
		for j := 0; j < game.NPegs; j++ {
			if ui.currentPeg == rune(j) {
				style = currentPegStyle
			} else {
				style = normalPegStyle
			}
			if len(ui.game.Pegs[j]) > i {
				s += style.Render(string(ui.game.Pegs[j][i] + '0') + " ")
			} else {
				s += style.Render("  ")
			}
		}
		s += "\n"
	}
	s += "\n"
	// peg labels
	for i := 0; i < game.NPegs; i++ {
		if ui.currentPeg == rune(i) {
			style = currentPegStyle.Bold(true)
		} else {
			style = normalPegStyle.Bold(true)
		}
		s += style.Render(string('A' + rune(i)) + " ")
		style.Bold(false)
	}
	s += "\n"
	return s
}

func (ui *UI) View() string {
	s := "\n"
	if ui.done {
		s += "You win!\n"
	} else if ui.quitting {
		return "Goodbye!\n"
	} else if ui.rules {
		s += ViewRules() + "\n"
	} else {
		s += ui.viewBoard() + "\n"
		if ui.err != nil {
			s += errorStyle.Render("Error: ")
			s += ui.err.Error() + "\n"
		} else {
			s += "\n\n"
		}
	}

	// wrap the rendered string to the width of the terminal
	s = wordwrap.String(s, ui.screenWidth - 1)

	helpView := ui.help.View(keys)

	// add padding to the bottom of the screen for the help view
	height := ui.screenHeight - strings.Count(s, "\n") - strings.Count(helpView, "\n")
	
	return s + strings.Repeat("\n", height) + helpView
}
