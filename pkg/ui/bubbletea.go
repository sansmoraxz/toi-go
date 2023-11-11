package ui

import (
	"errors"
	"strings"

	"github.com/sansmoraxz/toi-go/pkg/game"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UI struct {
	game *game.Hanoi
	currentPeg rune

	err error
	done bool
	quitting bool

	help help.Model
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
		"a": k.A,
		"d": k.D,
		"←": k.Left,
		"→": k.Right,
	}
}

var keys = keyMap{
	Quit:  key.NewBinding(
		key.WithKeys("q", "esc"),
		key.WithHelp("q", "Quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Toggle help"),
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
		key.WithKeys("r"),
		key.WithHelp("r", "Reload"),
	),
	A: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "Move disk left"),
	),
	D: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "Move disk right"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.R, k.Quit, k.Help}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.R, k.Quit, k.Help},			// second column
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

		help: help.New(),
	}
}

func (ui *UI) Start() error {
	p := tea.NewProgram(ui)
	_, err := p.Run()
	return err
}

func (ui *UI) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate its view as needed.
		ui.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			ui.quitting = true
			return ui, tea.Quit
		case key.Matches(msg, keys.Help):
			ui.help.ShowAll = !ui.help.ShowAll
		// arrow keys to move the cursor
		case key.Matches(msg, keys.Left):
			ui.currentPeg = game.PrevPeg(ui.currentPeg)
		case key.Matches(msg, keys.Right):
			ui.currentPeg = game.NextPeg(ui.currentPeg)
		case key.Matches(msg, keys.R):
			ui.game.Reset()
			ui.err = errors.New("game reset")
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

	return ui, nil
}

func viewRules() string {
	return rulesStyle.Render("Rules:\n" +
		"1. Only one disk can be moved at a time.\n" +
		"2. Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack.\n" +
		"3. No disk may be placed on top of a smaller disk.\n" +
		"4. Fill the last peg with all the disks to win.\n" + 
		"5. You can only move adjacent disks.\n")
}

func (ui *UI) viewBoard() string {
	s := ""
	for i := 0; i < game.NPegs; i++ {
		if ui.currentPeg == rune(i) {
			s += currentPegStyle.Render(string(rune(i) + 'A'))
		} else {
			s += normalPegStyle.Render(string(rune(i) + 'A'))
		}
		s += ": "
		for j := 0; j < len(ui.game.Pegs[i]); j++ {
			s += string(ui.game.Pegs[i][j] + '0') + " "
		}
		s += "\n"
	}
	return s
}

func (ui *UI) View() string {
	if ui.done {
		return "You win!\n"
	}
	if ui.quitting {
		return "Goodbye!\n"
	}
	s := ""
	s += viewRules() + "\n"
	s += ui.viewBoard() + "\n"
	if ui.err != nil {
		s += errorStyle.Render("Error: ")
		s += ui.err.Error() + "\n"
	}

	helpView := ui.help.View(keys)
	height := strings.Count(s, "\n") - strings.Count(helpView, "\n")
	
	return s + strings.Repeat("\n", height) + helpView
}
