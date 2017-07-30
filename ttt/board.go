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

type Victory struct {
	mark Marker

	start int
	end   int
}

func (b *Board) CheckForWinner(lastMoveCell int) *Victory {
	row, col := b.cell2RowCol(lastMoveCell)
	mark := b.cells[lastMoveCell]

	//vertical
	vic := b.check(b.generateCellList(row-b.victoryNumber+1, col, 1, 0), mark)
	if vic != nil {
		return vic
	}

	//horizontal
	vic = b.check(b.generateCellList(row, col-b.victoryNumber+1, 0, 1), mark)
	if vic != nil {
		return vic
	}

	//diagonal down
	vic = b.check(b.generateCellList(row-b.victoryNumber+1, col-b.victoryNumber+1, 1, 1), mark)
	if vic != nil {
		return vic
	}

	//diagonal up
	vic = b.check(b.generateCellList(row+b.victoryNumber-1, col-b.victoryNumber+1, -1, 1), mark)
	return vic
}

func (b *Board) generateCellList(startRow int, startCol int, dirRow int, dirCol int) []int {
	results := []int{}
	for i := 0; i < b.victoryNumber*2-1; i++ {
		r, c := startRow+(i*dirRow), startCol+(i*dirCol)
		if r < 0 || c < 0 || r >= b.Rows || c >= b.Cols {
			continue
		}
		results = append(results, b.rowCol2Cell(r, c))
	}
	return results
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
