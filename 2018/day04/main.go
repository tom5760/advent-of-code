package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

type EventKind int

const (
	BeginShift EventKind = iota
	FallAsleep
	WakeUp
)

type Event struct {
	Timestamp time.Time
	ID        int
	Kind      EventKind
}

func main() {
	events := sortEvents(readInput(os.Stdin))

	sleepMinutes, sleepTotals := sleep(events)

	var maxID, maxMinute int

	maxID, maxMinute = mostMinutes(sleepMinutes, sleepTotals)
	log.Printf("(part 1) most minutes: id * minute = %d * %d = %d", maxID, maxMinute, maxID*maxMinute)

	maxID, maxMinute = asleepOnSameMinute(sleepMinutes)
	log.Printf("(part 2) same minute: id * minute = %d * %d = %d", maxID, maxMinute, maxID*maxMinute)
}

const (
	timestampFormat  = "[%d-%d-%d %d:%d]"
	guardShiftFormat = "Guard #%d begins shift"
)

func readInput(r io.Reader) []Event {
	var events []Event

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var year, month, day, hour, minute int

		_, err := fmt.Sscanf(scanner.Text(), timestampFormat,
			&year, &month, &day, &hour, &minute)
		if err != nil {
			log.Fatalln("failed to parse timestamp:", err)
			return nil
		}

		evt := Event{Timestamp: time.Date(
			year, time.Month(month), day, hour, minute, 0, 0, time.UTC)}

		str := scanner.Text()[19:]

		switch str {
		case "wakes up":
			evt.Kind = WakeUp

		case "falls asleep":
			evt.Kind = FallAsleep

		default:
			evt.Kind = BeginShift
			_, err := fmt.Sscanf(str, guardShiftFormat, &evt.ID)
			if err != nil {
				log.Fatalln("failed to scan event id:", err)
				return nil
			}
		}

		events = append(events, evt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("error reading input:", err)
		return nil
	}

	return events
}

func sortEvents(events []Event) []Event {
	sort.Slice(events, func(i, j int) bool {
		iTS := events[i].Timestamp
		jTS := events[j].Timestamp
		return iTS.Before(jTS)
	})

	return events
}

func sleep(events []Event) (sleepMinutes [][]int, sleepTotals map[int]int) {
	// 60 minutes in the midnight hour
	sleepMinutes = make([][]int, 60)
	sleepTotals = make(map[int]int)

	var lastID, sleepStart int
	for _, evt := range events {
		switch evt.Kind {
		case BeginShift:
			lastID = evt.ID

		case FallAsleep:
			sleepStart = evt.Timestamp.Minute()

		case WakeUp:
			sleepEnd := evt.Timestamp.Minute()
			sleepTotals[lastID] += sleepEnd - sleepStart

			for i := sleepStart; i < sleepEnd; i++ {
				sleepMinutes[i] = append(sleepMinutes[i], lastID)
			}
		}
	}

	return sleepMinutes, sleepTotals
}

func mostMinutes(sleepMinutes [][]int, sleepTotals map[int]int) (id, minute int) {
	var maxID, maxSleep int
	for id, total := range sleepTotals {
		if total > maxSleep {
			maxSleep = total
			maxID = id
		}
	}

	var maxMinute, maxMinuteCount int
	for i := 0; i < 60; i++ {
		var minuteCount int
		for _, id := range sleepMinutes[i] {
			if id == maxID {
				minuteCount++
			}
		}

		if minuteCount > maxMinuteCount {
			maxMinuteCount = minuteCount
			maxMinute = i
		}
	}

	return maxID, maxMinute
}

func asleepOnSameMinute(sleepMinutes [][]int) (id, minute int) {
	var maxID, maxMinute, maxMinuteCount int
	for i := 0; i < 60; i++ {
		sleepCounts := make(map[int]int)

		for _, id := range sleepMinutes[i] {
			sleepCounts[id]++
		}

		for id, count := range sleepCounts {
			if count > maxMinuteCount {
				maxID = id
				maxMinute = i
				maxMinuteCount = count
			}
		}
	}

	return maxID, maxMinute
}
