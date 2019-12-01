package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
)

const ScoringMultiple = 23

type Circle struct {
	circle list.List
	cur    *list.Element
}

func (c *Circle) Place(marble uint64) {
	left := c.clockwise(1)
	c.cur = c.circle.InsertAfter(marble, left)
}

func (c *Circle) Score(marble uint64) uint64 {
	removeEl := c.counter(7)
	c.cur = c.next(removeEl)

	removeValue := c.circle.Remove(removeEl).(uint64)

	return marble + removeValue
}

func (c *Circle) next(el *list.Element) *list.Element {
	el = el.Next()
	if el == nil {
		el = c.circle.Front()
	}
	return el
}

func (c *Circle) prev(el *list.Element) *list.Element {
	el = el.Prev()
	if el == nil {
		el = c.circle.Back()
	}
	return el
}

func (c *Circle) clockwise(count int) *list.Element {
	el := c.cur
	for i := 0; i < count; i++ {
		el = c.next(el)
	}
	return el
}

func (c *Circle) counter(count int) *list.Element {
	el := c.cur
	for i := 0; i < count; i++ {
		el = c.prev(el)
	}
	return el
}

func (c *Circle) pushBack(marble uint64) *list.Element {
	return c.circle.PushBack(marble)
}

type Game struct {
	PlayerCount uint64
	LastMarble  uint64
}

func (g *Game) Play() uint64 {
	scores := make(map[uint64]uint64)
	var circle Circle

	// Place the first two marbles in the circle
	circle.pushBack(0)
	circle.cur = circle.pushBack(1)

	for i := uint64(2); i <= g.LastMarble; i++ {
		if i%ScoringMultiple == 0 {
			player := i % g.PlayerCount
			scores[player] += circle.Score(i)

		} else {
			circle.Place(i)
		}
	}

	var highScore uint64
	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}

func main() {
	games := readInput(os.Stdin)

	for _, game := range games {
		log.Println("(part 1) high score:", game.Play())

		game.LastMarble *= 100
		log.Println("(part 2) x100 high score:", game.Play(), "\n")
	}
}

const inputFormat = "%d players; last marble is worth %d points"

func readInput(r io.Reader) []Game {
	var games []Game

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var game Game

		if _, err := fmt.Sscanf(scanner.Text(), inputFormat, &game.PlayerCount, &game.LastMarble); err != nil {
			log.Fatalln("failed to parse input:", err)
			return nil
		}

		games = append(games, game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return games
}
