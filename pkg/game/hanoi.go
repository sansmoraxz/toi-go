package game

import (
	"fmt"
	"math/rand"
)

const NPegs = 3 // number of pegs
const NDsks = 5  // number of disks

type Hanoi struct {
	Pegs [][]rune
	resetPegs [][]rune
}

func NewHanoi() *Hanoi {

	// create pegs
	pegs := make([][]rune, NPegs)
	for i := 0; i < NPegs; i++ {
		// fill pegs with disks only if it is the first peg
		if i == 0 {
			pegs[i] = make([]rune, NDsks)
			// randomize the order of disks
			sl := rand.Perm(NDsks)
			for j := 0; j < NDsks; j++ {
				pegs[i][j] = rune(sl[j]) + 1
			}
		} else {
			pegs[i] = make([]rune, 0)
		}
	}

	// for resetting the game
	var resetPegs [][]rune
	for i := 0; i < NPegs; i++ {
		resetPegs = append(resetPegs, make([]rune, len(pegs[i])))
		copy(resetPegs[i], pegs[i])
	}

	return &Hanoi{
		Pegs: pegs,
		resetPegs: resetPegs,
	}
}

func (h *Hanoi) MoveDisk(from, to rune) error {
	if len(h.Pegs[from]) == 0 {
		return fmt.Errorf("no disk to move on peg %d", from)
	}
	disk := h.Pegs[from][len(h.Pegs[from])-1]
	if len(h.Pegs[to]) > 0 && h.Pegs[to][len(h.Pegs[to])-1] < disk {
		return fmt.Errorf("disk %d is larger than the top disk on peg %d", disk, to)
	}
	h.Pegs[from] = h.Pegs[from][:len(h.Pegs[from])-1]
	h.Pegs[to] = append(h.Pegs[to], disk)
	return nil
}


// if last peg is full, then game is finished
func (h *Hanoi) IsFinished() bool {
	return len(h.Pegs[0]) == 0 && len(h.Pegs[NPegs-1]) == NDsks
}

func (h *Hanoi) Reset() {
	for i := 0; i < NPegs; i++ {
		h.Pegs[i] = make([]rune, len(h.resetPegs[i]))
		copy(h.Pegs[i], h.resetPegs[i])
	}
}

func roundOffPeg(i rune) rune {
	if i < 0 {
		i = NPegs - 1
	} else if i >= NPegs {
		i = 0
	}
	return i
}


func NextPeg(currentPeg rune) rune {
	return roundOffPeg(currentPeg + 1)
}

func PrevPeg(currentPeg rune) rune {
	return roundOffPeg(currentPeg - 1)
}
