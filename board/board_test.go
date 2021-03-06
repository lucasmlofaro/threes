package board

import (
	. "Threes/tile"
	"fmt"
	"testing"
)

func TestBoard(t *testing.T) {
	var board Board

	if len(board) != 4 {
		t.Error("Board should have 4 rows")
	}

	for _, row := range board {
		if len(row) != 4 {
			t.Error("Row should have 4 entires")
		}
	}
}

func TestBoard_String(t *testing.T) {
	var board Board
	boardString := `|     0 |     0 |     0 |     0 |
|     0 |     0 |     0 |     0 |
|     0 |     0 |     0 |     0 |
|     0 |     0 |     0 |     0 |`

	expected := boardString
	actual := fmt.Sprint(board)
	if actual != expected {
		t.Errorf("New board printed incorecctly. Got:\n%v\nWanted:\n%v\n", actual, expected)
	}
}

func TestRow_String(t *testing.T) {
	var row Row
	rowString := `|     0 |     0 |     0 |     0 |`

	expected := rowString
	actual := fmt.Sprint(row)
	if actual != expected {
		t.Errorf("New row printed incorecctly. Got:\n%v\nWanted:\n%v\n", actual, expected)
	}
}

func TestBoard_AddTile(t *testing.T) {
	var tile Tile
	tile.SetValue(3)

	var board Board
	index := Index{0, 0}

	board.AddTileAt(tile, index)

	if board[index.Row][index.Column].IsEmpty() {
		t.Error("Tile should not be empty")
	}
}

func TestBoard_Slide(t *testing.T) {
	var board Board
	var indices []Index

	var t1, t2, t3, t4, t5, t6, t7 Tile

	t1.SetValue(3)
	board.AddTileAt(t1, Index{0, 0})
	t2.SetValue(3)
	board.AddTileAt(t2, Index{0, 1})
	t3.SetValue(3)
	board.AddTileAt(t3, Index{1, 1})
	t4.SetValue(3)
	board.AddTileAt(t4, Index{2, 1})
	t5.SetValue(3)
	board.AddTileAt(t5, Index{0, 3})
	t6.SetValue(3)
	board.AddTileAt(t6, Index{1, 3})
	t7.SetValue(3)
	board.AddTileAt(t7, Index{1, 0})
	fmt.Println(board)
	fmt.Println()
	indices, _ = board.Slide(Up)
	fmt.Println(indices)
	if len(indices) != 3 {
		t.Errorf("Incorrect number of next indices")
	}
	indices, _ = board.Slide(Right)
	if len(indices) != 2 {
		t.Errorf("Incorrect number of next indices")
	}
	indices, _ = board.Slide(Right)
	if len(indices) != 2 {
		t.Errorf("Incorrect number of next indices")
	}
	fmt.Println(board)

	expected := `|     0 |     0 |     6 |    12 |
|     0 |     0 |     0 |     3 |
|     0 |     0 |     0 |     0 |
|     0 |     0 |     0 |     0 |`

	if board.String() != expected {
		t.Errorf("Incorrect board configuration: expected\n%s\nactual\n%s\n", expected, board)
	}

	board = Board{}
	fmt.Println()
	t1.SetValue(6)
	board.AddTileAt(t1, Index{0, 0})
	t2.SetValue(6)
	board.AddTileAt(t2, Index{0, 1})
	t3.SetValue(6)
	board.AddTileAt(t3, Index{1, 1})
	t4.SetValue(6)
	board.AddTileAt(t4, Index{2, 1})
	t5.SetValue(2)
	board.AddTileAt(t5, Index{0, 3})
	t6.SetValue(1)
	board.AddTileAt(t6, Index{1, 3})
	t7.SetValue(3)
	board.AddTileAt(t7, Index{1, 0})
	fmt.Println(board)
	fmt.Println()
	indices, _ = board.Slide(Left)
	if len(indices) != 3 {
		t.Errorf("Incorrect number of next indices")
	}
	indices, _ = board.Slide(Down)
	if len(indices) != 3 {
		t.Errorf("Incorrect number of next indices")
	}
	indices, _ = board.Slide(Down)
	if len(indices) != 2 {
		t.Errorf("Incorrect number of next indices")
	}
	indices, _ = board.Slide(Down)
	if len(indices) != 1 {
		t.Errorf("Incorrect number of next indices")
	}
	fmt.Println(board)
	indices, _ = board.Slide(Left)
	fmt.Println(indices)
	if len(indices) != 1 {
		t.Errorf("Incorrect number of next indices")
	}
	fmt.Println(board)

	expected = `|     0 |     0 |     0 |     0 |
|    12 |     0 |     0 |     0 |
|     3 |     0 |     0 |     0 |
|    12 |     3 |     0 |     0 |`

	if board.String() != expected {
		t.Errorf("Incorrect board configuration: expected\n%s\nactual\n%s\n", expected, board)
	}
}

func TestBoard_Score(t *testing.T) {
	var board Board

	var t1, t2, t3, t4, t5, t6, t7 Tile

	t1.SetValue(3)
	board.AddTileAt(t1, Index{0, 0})
	t2.SetValue(6)
	board.AddTileAt(t2, Index{0, 1})
	t3.SetValue(3)
	board.AddTileAt(t3, Index{1, 1})
	t4.SetValue(24)
	board.AddTileAt(t4, Index{2, 1})
	t5.SetValue(6)
	board.AddTileAt(t5, Index{0, 3})
	t6.SetValue(3)
	board.AddTileAt(t6, Index{1, 3})
	t7.SetValue(12)
	board.AddTileAt(t7, Index{1, 0})
	fmt.Println(board)

	var expected float64 = 135
	if board.Score() != expected {
		t.Errorf("Score is %v, but should be %v", board.Score(), expected)
	}
}
