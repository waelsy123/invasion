package earth

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

const (
	NORTH = "north"
	EAST  = "east"
	SOUTH = "south"
	WEST  = "west"
)

type Connection [3]string
type Board struct {
	connections    []Connection
	alienLocations map[string][]int
}

func (board *Board) Init(filename string, n int) {
	board.connections = readConnections(filename)
	board.distributeAliens(n)
}

func (board *Board) GetConnections() []Connection {
	return board.connections
}

func (board *Board) GetAlienLocations() map[string][]int {
	return board.alienLocations
}

func (board *Board) Print() {
	currentCity := ""
	for _, connection := range board.connections {
		mainCity := connection[0]
		desCity := connection[1]
		direction := connection[2]

		if mainCity != currentCity {
			currentCity = mainCity
			fmt.Printf("\n%s", mainCity)

		}
		fmt.Printf(" %s=%s", direction, desCity)
	}
}

func (board *Board) distributeAliens(n int) {
	alienLocations := map[string][]int{}

	for _, item := range board.connections {
		alienLocations[item[0]] = make([]int, 0)
		alienLocations[item[1]] = make([]int, 0)
	}

	var cities []string
	for name := range alienLocations {
		cities = append(cities, name)
	}

	for i := 0; i < n; i++ {
		randomIdx := rand.Intn(len(cities))
		randomCity := cities[randomIdx]

		alienLocations[randomCity] = append(alienLocations[randomCity], i)
	}

	board.alienLocations = alienLocations
}

func (board *Board) MovingPhase() {
	connectionsByCity := getConnectionsByCity(board.connections)

	newAlienLocations := map[string][]int{}

	for city, aliensInCity := range board.alienLocations {
		cityConnections, isConnectionExists := connectionsByCity[city]

		if len(aliensInCity) >= 2 {
			log.Println("two aliens in a city, forgot to destroy?")
		}

		if len(aliensInCity) > 0 && isConnectionExists {
			randomIdx := rand.Intn(len(cityConnections))
			connectionIdx := connectionsByCity[city][randomIdx]

			desCity := board.connections[connectionIdx].getDestinationCity(city)

			newAlienLocations[desCity] = append(newAlienLocations[desCity], aliensInCity[0])
		}
	}

	board.alienLocations = newAlienLocations
}

func (board *Board) DestoryPhase() {
	citiesToBeDestoryed := map[string]bool{}

	for city, aliens := range board.alienLocations {
		if len(aliens) > 1 {
			log.Printf("The city %s will be destroyed by aliens: %v\n", city, aliens)

			// destory city
			delete(board.alienLocations, city)
			citiesToBeDestoryed[city] = true
		}
	}

	// destory connection to city
	newConnections := []Connection{}
	for _, connection := range board.connections {
		cityA := connection[0]
		cityB := connection[1]

		if !citiesToBeDestoryed[cityB] && !citiesToBeDestoryed[cityA] {
			newConnections = append(newConnections, connection)
		}
	}

	board.connections = newConnections
}

func (c *Connection) getDestinationCity(sourceCity string) string {
	if c[0] == sourceCity {
		return c[1]
	}

	return c[0]
}

func getConnectionsByCity(connections []Connection) map[string][]int {
	cityConnections := map[string][]int{}
	for idx, connection := range connections {
		cityConnections[connection[0]] = append(cityConnections[connection[0]], idx)
		cityConnections[connection[1]] = append(cityConnections[connection[1]], idx)
	}

	return cityConnections
}

func readConnections(filename string) []Connection {
	var (
		text        []string
		connections []Connection
	)

	file, err := os.Open(filename)

	if err != nil {
		log.Panicf("error opening file: %v\n", err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	if len(text) == 0 {
		log.Panic("empty file.. \n")
	}

	for _, line := range text {
		lineparts := strings.Split(line, " ")
		cityName := lineparts[0]

		connectionsLine := lineparts[1:]
		for _, connectionString := range connectionsLine {
			connectionParts := strings.Split(connectionString, "=")
			direction := connectionParts[0]
			neighbourCity := connectionParts[1]

			if direction != NORTH && direction != EAST && direction != SOUTH && direction != WEST {
				log.Panic("connection should be one of the following values: north south east west")
			}

			connection := Connection{cityName, neighbourCity, direction}
			connections = append(connections, connection)

		}
	}

	return connections
}
