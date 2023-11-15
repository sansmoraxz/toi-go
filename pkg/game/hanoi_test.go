package game

import (
	"reflect"
	"testing"
)


func TestNewHanoi(t *testing.T) {
	h := NewHanoi()

	// check that all pegs except the first one are empty
	for i := 1; i < NPegs; i++ {
		if len(h.Pegs[i]) != 0 {
			t.Errorf("Peg %d is not empty: %v", i, h.Pegs[i])
		}
	}

	// check that the first peg has all the disks
	if len(h.Pegs[0]) != NDsks {
		t.Errorf("Peg 0 does not have %d disks: %v", NDsks, h.Pegs[0])
	}

	// check that resetPegs is a copy of Pegs
	for i := 0; i < NPegs; i++ {
		if !reflect.DeepEqual(h.Pegs[i], h.resetPegs[i]) {
			t.Errorf("resetPegs is not a copy of Pegs: %v != %v", h.Pegs, h.resetPegs)
			break
		}
	}
}



func newTestHanoi(NPegs int, NDsks int) *Hanoi {
	// create pegs


	pegs := make([][]rune, NPegs)
	for i := 0; i < NPegs; i++ {
		// fill pegs with disks only if it is the first peg
		if i == 0 {
			pegs[i] = make([]rune, NDsks)
			for j := 0; j < NDsks; j++ {
				pegs[i][j] = rune(j) + 1
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



func TestMoveDisk(t *testing.T) {
	const NPegs = 3 // number of pegs
	const NDsks = 3  // number of disks


	h := newTestHanoi(NPegs, NDsks)

	// move disk from peg 0 to peg 1
	err := h.MoveDisk(0, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(h.Pegs[0]) != NDsks-1 {
		t.Errorf("disk not removed from peg %s: %v", string('A'), h.Pegs[0])
	}
	if len(h.Pegs[1]) != 1 {
		t.Errorf("disk not added to peg %s: %v", string('B'), h.Pegs[1])
	}

	// move disk from peg 0 to peg 2
	err = h.MoveDisk(0, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(h.Pegs[0]) != NDsks-2 {
		t.Errorf("disk not removed from peg %s: %v", string('A'), h.Pegs[0])
	}
	if len(h.Pegs[2]) != 1 || h.Pegs[2][0] != NDsks-1 {
		t.Errorf("disk not added to peg %s: %v", string('C'), h.Pegs[2])
	}

	// move disk from peg 0 to peg 2
	err = h.MoveDisk(0, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(h.Pegs[0]) != 0 {
		t.Errorf("disk not removed from peg %s: %v", string('A'), h.Pegs[0])
	}
	if len(h.Pegs[2]) != 2 || h.Pegs[2][0] != NDsks-1 || h.Pegs[2][1] != NDsks-2 {
		t.Errorf("disk not added to peg %s: %v", string('C'), h.Pegs[2])
	}

	// try to move disk from empty peg
	err = h.MoveDisk(0, 2)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
	if len(h.Pegs[0]) != 0 {
		t.Errorf("disk removed from empty peg %s: %v", string('B'), h.Pegs[1])
	}
	if len(h.Pegs[2]) != 2 {
		t.Errorf("disk added to peg %s: %v", string('C'), h.Pegs[2])
	}

	// try to move larger disk onto smaller disk
	err = h.MoveDisk(1, 2)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
	if len(h.Pegs[1]) != 1 {
		t.Errorf("disk removed from peg %s: %v", string('A'), h.Pegs[0])
	}
	if len(h.Pegs[2]) != 2 {
		t.Errorf("disk added to peg %s: %v", string('B'), h.Pegs[1])
	}
}
