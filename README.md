# TOI

This game is a terminal implementation of the classic [Tower of Hanoi](https://en.wikipedia.org/wiki/Tower_of_Hanoi) game.

## Set Up

You need to have Go installed on your system. If you don't have Go installed, you can download it from the [official website](https://golang.org/dl/).


Clone the project with the following command:

```bash
git clone git@github.com:sansmoraxz/toi-go.git
```

# Running the Game

To start the game, navigate to the directory where you cloned the project and run the following command:

```bash
go run cmd/toi/main.go
```

The game will start and you will see the initial state of the game in your terminal.

The goal of the game is to move all the disks from the leftmost peg to the rightmost peg. You can only move one disk at a time and you cannot place a larger disk on top of a smaller disk.

# Navigation

You can move the cursor around the pegs using the arrow keys. The current position of the cursor is highlighted in blue.

To move the disks use the `A` or `D` keys. `A` moves the disk to the left and `D` moves the disk to the right.


## Resetting the Game

If you want to restart a new game, you can reset the game by pressing `r`.

## Quitting the Game

To quit the game, you can press `q`.

## Enjoy the Game!

We hope you enjoy playing Tower of Hanoi! If you encounter any issues or have any suggestions, please feel free to open an issue.
