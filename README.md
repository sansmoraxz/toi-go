# Tower of Hanoi

This is a simple implementation of the Tower of Hanoi game using Go and the Bubbletea library for the terminal-based user interface.

## Installation

To install the game, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the official website.

Once you have Go installed, you can install the game by running the following command:

```bash
go get github.com/yourusername/tower-of-hanoi
```

Replace `yourusername` with your actual GitHub username.

## How to Play

To start the game, navigate to the directory where you installed the game and run the following command:

```bash
go run cmd/tower-of-hanoi/main.go
```

The game will start and you will see the initial state of the game in your terminal.

The goal of the game is to move all the disks from the leftmost peg to the rightmost peg. You can only move one disk at a time and you cannot place a larger disk on top of a smaller disk.

To move a disk, you can use the arrow keys to select the disk you want to move and then press `a` or `d` to move the disk to the left or right peg respectively.

## Resetting the Game

If you want to start a new game, you can reset the game by entering `r`.

## Quitting the Game

To quit the game, you can enter `q`.

## Enjoy the Game!

We hope you enjoy playing Tower of Hanoi! If you encounter any issues or have any suggestions, please feel free to open an issue.
