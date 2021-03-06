# Threes

An AI for the puzzle game Threes!

Threes! is a game in which users combine numeric tiles to generate the most possible points. This project attempts to create an artificially intelligent agent using exectimax search.

## Run the Project
From the root of the repository, execute in bash:
```bash
$ go run main.go
```

## Test the Code
Each package has its own tests. Package tests can be run individually with:
```bash
$ go test <package-name>
```

Alternatively, all tests can be run with:
```bash
$ go test ./...
```

## Board Repreesentation
The board is made up of a 4-by-4 grid of tiles. Tiles will slide one space at a time and combine with each other occording to their respective values. The board is scored according to point assignment system that heavily favors larger tiles. Each tile is worth 3^_n_ points, where _n_ is the number times the tile has been merged.

### Implementation
There are 4 allowable moves: `Up`, `Left`, `Down`, and `Right`. Since the board is a perfect square, swiping in any direction is equivalent to first rotating the board by 0, 90, 180, or 270 degress and then swiping up. This coordinate tranformation can be found in the `Board`'s unexported `rotate()` method.

## Tile Repreesentation
Tiles can have a value of 1, 2, or any multiple of 3 (up to the maximum tile). Tiles with value 1 can only merge with tiles of value 2, while tiles with value that is a multiple of 3 can only be merged with other tiles of equal value.

### Generating tiles
One new tile is generated after each move. Tiles are usually generated by a `SimpleTileGenerator` which distributes tiles from a known stack wth a ranomized order. The stack consists of four tiles of value 1, four tiles of value 2, and four tiles of value 3.
Occasionally, bonus tiles will be generated via the `BonusTileGenerator`. Bonus tiles are only available once a tile with a value of 48 appears on the board. Bonus tiles start with a value that is 1/8 of the maximum tile value on the board.
