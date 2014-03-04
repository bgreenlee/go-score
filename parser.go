package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"footle.org/go-score/board"
)

func parseParams(paramList string) []string {
	// remove leading and trailing brackets, then split on ][
	return strings.Split(paramList[1:len(paramList)-1], "][")
}

func main() {
	var theBoard board.Board

	// check commandline args
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <board.sgf>\n", os.Args[0])
	}

	// read in sgf file
	inputFilename := os.Args[1]
	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 2 {
			// first two characters are the command, the remaining are the parameters
			cmd := line[:2]
			// many commands use the second char to represent the color, so grab that now
			color := board.Black
			if cmd[1:] == "W" {
				color = board.White
			}
			params := parseParams(line[2:])
			// switch on the first character
			// this would probably not work on the full SGF spec, but we're just dealing with a
			// greatly simplified version
			switch cmd[:1] {
			case "S": // size command; initialize board
				boardSize, err := strconv.Atoi(params[0])
				if err != nil {
					log.Fatal("Error parsing board size: ", err)
				}
				theBoard = board.New(boardSize)
			case "A": // initial stone placement
				for _, coord := range params {
					theBoard.Set(coord, color)
				}
			case ";": // move
				coord := params[0]
				if len(coord) > 0 {
					theBoard.Set(coord, color)
				}
			case "T": // end-of-game territories list
				for _, coord := range params {
					var updatedPoint byte
					switch theBoard.At(coord) {
					case board.White:
						updatedPoint = board.DeadWhite
					case board.Black:
						updatedPoint = board.DeadBlack
					default:
						updatedPoint = board.Empty
					}
					theBoard.Set(coord, updatedPoint)
				}
			}
		}
	}

	fmt.Println(theBoard)
}
