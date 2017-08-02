package ttt

import (
	"bytes"
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
		return fmt.Errorf("cell already filled")
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

type Victory struct {
	mark Marker

	start int
	end   int
}

// Represents a direction to move on the tic tac toe board
type direction struct {
	row int
	col int
}

var (
	HORZ   direction = direction{0, 1}  // horizontal
	VERT             = direction{1, 0}  // vertical
	DIAG_U           = direction{1, 1}  // diagonal up
	DIAG_D           = direction{-1, 1} // diagonal down
)

var directions []direction = []direction{HORZ, VERT, DIAG_U, DIAG_D}

func (b *Board) CheckForWinner(lastMoveCell int) *Victory {
	row, col := b.cell2RowCol(lastMoveCell)
	mark := b.cells[lastMoveCell]
	var vic *Victory
	for _, dir := range directions {
		vic = b.check(b.generateCellList(row, col, dir), mark)
		if vic != nil {
			return vic
		}
	}
	return vic
}

// generates the possible places for a win centered around midRow, midCol in the given direction
func (b *Board) generateCellList(midRow int, midCol int, dir direction) (locs []int) {
	startRow, startCol := midRow-(dir.row*(b.victoryNumber-1)), midCol-(dir.col*(b.victoryNumber-1))

	for i := 0; i < b.victoryNumber*2-1; i++ {
		r, c := startRow+(i*dir.row), startCol+(i*dir.col)
		if r < 0 || c < 0 || r >= b.Rows || c >= b.Cols {
			continue
		}
		locs = append(locs, b.rowCol2Cell(r, c))
	}
	return
}

func (b *Board) check(locs []int, mark Marker) *Victory {

	numAdj := 0
	for i, cell := range locs {
		m := b.cells[cell]
		if m != mark {
			numAdj = 0
		}

		if m == mark {
			numAdj += 1
			if numAdj >= b.victoryNumber {
				return &Victory{
					mark:  mark,
					start: locs[i-b.victoryNumber+1],
					end:   cell,
				}
			}
		}
	}

	return nil
}

func (b *Board) Find(mark Marker) (cells []int) {
	for cell, m := range b.cells {
		if m == mark {
			cells = append(cells, cell)
		}
	}
	return
}

func (b *Board) IsFull() bool {
	return len(b.Find(BLANK)) == 0
}

func (b Board) String() (s string) {
	var buf bytes.Buffer
	sep := "|"

	for i, m := range b.cells {
		fmt.Fprint(&buf, m)
		if i%b.Cols == b.Cols-1 {
			fmt.Fprint(&buf, "\n")
		} else {
			fmt.Fprint(&buf, sep)
		}
	}

	return buf.String()
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
