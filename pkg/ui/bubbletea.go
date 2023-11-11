package ui

import (
	"errors"

	"github.com/sansmoraxz/toi-go/pkg/game"

	tea "github.com/charmbracelet/bubbletea"
)

type UI struct {
	game *game.Hanoi
	currentPeg rune
	err error
	done bool
}

func NewUI() *UI {
	hanoi := game.NewHanoi()
	return &UI{
		game: hanoi,
		currentPeg: 0,
		err: nil,
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
	case tea.KeyMsg:
		switch msg.String() {
		// q or ctrl+c to quit
		case "q", "ctrl+c":
			return ui, tea.Quit
		// arrow keys to move the cursor
		case "left":
			ui.currentPeg = game.PrevPeg(ui.currentPeg)
		case "right":
			ui.currentPeg = game.NextPeg(ui.currentPeg)
		case "r":
			ui.game.Reset()
			ui.err = errors.New("game reset")
		// a or d to move disk
		case "a":
			newPeg := game.PrevPeg(ui.currentPeg)
			ui.err = ui.game.MoveDisk(ui.currentPeg, newPeg)
			if ui.err != nil {
				ui.currentPeg = newPeg
			}
		case "d":
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

func rules() string {
	return "Rules:\n" +
		"1. Only one disk can be moved at a time.\n" +
		"2. Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack.\n" +
		"3. No disk may be placed on top of a smaller disk.\n" +
		"4. Fill the last peg with all the disks to win.\n"
}

func (ui *UI) View() string {
	if ui.done {
		return "You win!\n"
	}
	s := ""
	s += rules() + "\n"
	for i := 0; i < game.NPegs; i++ {
		s += "Peg " + string(rune(i + 'A')) + ": "
		for j := 0; j < len(ui.game.Pegs[i]); j++ {
			s += string(ui.game.Pegs[i][j] + '0') + " "
		}
		s += "\n"
	}
	s += "\n"
	s += "Current peg: " + string(ui.currentPeg + 'A') + "\n"
	if ui.err != nil {
		s += "Error: " + ui.err.Error() + "\n"
	}
	
	return s
}
