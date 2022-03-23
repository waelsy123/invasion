package earth

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	NORTH = "north"
	EAST  = "east"
	SOUTH = "south"
	WEST  = "west"
)

type City struct {
	name       string
	neighbours map[string]string
	aliens     int
}

type Board struct {
	dic         map[string]City
	connections [3]string // A - B - direction
}

func (board Board) cities() []string {
	var list []string

	for k := range board.dic {
		list = append(list, k)
	}

	return list
}

func CreateBoard(filename string, n int) (Board, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	board := make(Board)

	for _, line := range text {
		var city City

		lineparts := strings.Split(line, " ")

		city.name = lineparts[0]
		city.neighbours = make(map[string]string)

		connections := lineparts[1:]

		for _, connection := range connections {
			connectionParts := strings.Split(connection, "=")
			direction := connectionParts[0]
			neighbourCity := connectionParts[1]

			if direction != NORTH && direction != EAST && direction != SOUTH && direction != WEST {
				return nil, errors.New("connection should be one of the following values: north south east west")
			}

			city.neighbours[neighbourCity] = direction
		}

		board[city.name] = city
	}

	fmt.Println(board)

	// cityCount := len(board)
	// cities := board.cities()
	// log.Printf("c: %v\n", cities)

	// for i := 0; i < n; i++ {
	// 	randomIdx := rand.Intn(cityCount)
	// 	randomCity := cities[randomIdx]

	// 	log.Printf("randomCity: %v\n", randomCity)

	// }

	// fill with aliens
	return board, nil
}
