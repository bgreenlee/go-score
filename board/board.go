package board

import (
	"bytes"
	"fmt"
)

const (
	Empty = byte(0)
	Black = byte(1)
	White = byte(2)
)

type Board struct {
	size  int
	cells []byte
}

func New(size int) Board {
	board := Board{
		size:  size,
		cells: make([]byte, size*size),
	}

	return board
}

// convert a two-letter alpha coordinate to a board index
func (board Board) coordToIndex(coord string) int {
	xCoord := int(coord[0])
	yCoord := int(coord[1])
	// 97 is the ASCII value of "a"
	return (yCoord-97)*board.size + xCoord - 97
}

// set the value of a board at a given coordinate
// coordinate is given in SGF notation as a two-character string (xy)
func (board Board) Set(coord string, value byte) {
	board.cells[board.coordToIndex(coord)] = value
}

func (board Board) String() string {
	var buffer bytes.Buffer
	displayChars := [3]string{".", "B", "W"}
	for i := 0; i < len(board.cells); i++ {
		if i%board.size == 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteString(fmt.Sprintf("%s ", displayChars[board.cells[i]]))
	}
	buffer.WriteString("\n")
	return buffer.String()
}
