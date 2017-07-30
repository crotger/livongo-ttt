package ttt

import (
	"fmt"
)

type Marker int

const (
	BLANK Marker = iota
	MARK_X
	MARK_O
)

type Board struct {
	Rows          int
	Cols          int
	cells         []Marker
	victoryNumber int
}

func NewBoard(rows int, cols int, numInARow int) *Board {
	b := Board{}
	b.Rows = rows
	b.Cols = cols
	b.victoryNumber = numInARow

	b.cells = make([]Marker, b.Rows*b.Cols)

	return &b
}

func DefaultBoard() *Board {
	return NewBoard(3, 3, 3)
}

func (b *Board) Set(cell int, mark Marker) error {
	if cell > (b.Rows*b.Cols) || cell < 0 {
		return fmt.Errorf("cell out of bounds: %v", cell)
	}

	if b.cells[cell] != BLANK {
		return fmt.Errorf("cell %v already filled", cell)
	}

	b.cells[cell] = mark
	return nil
}

func (b *Board) Get(cell int) (Marker, error) {
	if cell >= (b.Rows*b.Cols) || cell < 0 {
		return BLANK, fmt.Errorf("cell out of bounds: %v", cell)
	}
	return b.cells[cell], nil
}

func (b *Board) cell2RowCol(cell int) (int, int) {
	return cell / b.Cols, cell % b.Cols
}

func (b *Board) rowCol2Cell(row int, col int) int {
	return row*b.Cols + col
}

func (m Marker) String() (s string) {
	switch m {
	case BLANK:
		s = " "
	case MARK_X:
		s = "X"
	case MARK_O:
		s = "O"
	default:
		s = "?"
	}
	return
}
