package board

import (
	"bytes"
	"fmt"
	"log"
)

// valid cell states
const (
	Empty     = byte(0)
	Black     = byte(1)
	White     = byte(2)
	DeadBlack = byte(3)
	DeadWhite = byte(4)
)

type Board struct {
	size   int
	points []byte
}

func New(size int) Board {
	if size > 26 {
		log.Fatalf("Sorry, can't do a %dx%[1]d board; the maximum is 26x26", size)
	}
	board := Board{
		size:   size,
		points: make([]byte, size*size),
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

// return the value of the board at the given coordinate
// coordinate is given in SGF notation as a two-character string (xy)
func (board Board) At(coord string) byte {
	return board.points[board.coordToIndex(coord)]
}

// set the value of a board at a given coordinate
// coordinate is given in SGF notation as a two-character string (xy)
func (board Board) Set(coord string, value byte) {
	board.points[board.coordToIndex(coord)] = value
}

func (board Board) String() string {
	var buffer bytes.Buffer
	displayChars := [5]string{".", "X", "O", "%", "0"}
	for i := 0; i < len(board.points); i++ {
		if i%board.size == 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteString(fmt.Sprintf("%s ", displayChars[board.points[i]]))
	}
	buffer.WriteString("\n")
	return buffer.String()
}
