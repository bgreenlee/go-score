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
		fmt.Printf("Usage: %s <board.sgf>\n", os.Args[0])
		os.Exit(1)
	}

	// read in sgf file
	inputFilename := os.Args[1]
	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 2 {
			cmd := line[:2]
			// many commands use the second char to represent the color, so try to grab that now
			color := board.Black
			if cmd[1:] == "W" {
				color = board.White
			}
			params := line[2:]
			// switch on the first character
			switch cmd[:1] {
			case "S": // size command; initialize board
				boardSize, err := strconv.Atoi(parseParams(params)[0])
				if err != nil {
					log.Fatal("Error parsing board size", err)
				}
				theBoard = board.New(boardSize)
			case "A": // initial stone placement
				coords := parseParams(params)
				for _, coord := range coords {
					theBoard.Set(coord, color)
				}
			case ";": // move
				coord := parseParams(params)[0]
				if len(coord) > 0 {
					theBoard.Set(coord, color)
				}
			}
		}
	}

	fmt.Println(theBoard)
}
