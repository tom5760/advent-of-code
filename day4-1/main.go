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

type eventKind int

const (
	beginShift eventKind = iota
	fallAsleep
	wakeUp
)

type event struct {
	timestamp time.Time
	id        int
	kind      eventKind
}

func main() {
	events := readInput(os.Stdin)

	sort.Slice(events, func(i, j int) bool {
		iTS := events[i].timestamp
		jTS := events[j].timestamp
		return iTS.Before(jTS)
	})

	var lastID, sleepStart int

	sleepTotals := make(map[int]int)

	// 60 minutes in the midnight hour
	sleepMinutes := make([][]int, 60)

	for _, evt := range events {
		switch evt.kind {
		case beginShift:
			lastID = evt.id

		case fallAsleep:
			sleepStart = evt.timestamp.Minute()

		case wakeUp:
			sleepEnd := evt.timestamp.Minute()
			total := sleepEnd - sleepStart
			sleepTotals[lastID] += total

			for i := sleepStart; i < sleepEnd; i++ {
				sleepMinutes[i] = append(sleepMinutes[i], lastID)
			}
		}
	}

	var maxID, maxSleep int
	for id, total := range sleepTotals {
		if total > maxSleep {
			maxSleep = total
			maxID = id
		}
	}

	log.Printf("guard #%d slept longest for %d minutes", maxID, maxSleep)

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

	log.Printf("guard #%d was asleep most during minute %d (%d days)",
		maxID, maxMinute, maxMinuteCount)

	log.Printf("total is %d * %d = %d", maxID, maxMinute, maxID*maxMinute)
}

func readInput(r io.Reader) []event {
	scanner := bufio.NewScanner(r)

	events := make([]event, 10)

	for scanner.Scan() {
		var year, month, day, hour, minute int

		if _, err := fmt.Sscanf(scanner.Text(), "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute); err != nil {
			log.Fatalln("failed to scan input line:", err)
			return nil
		}

		evt := event{
			timestamp: time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
		}

		str := scanner.Text()[19:]

		switch str {
		case "wakes up":
			evt.kind = wakeUp

		case "falls asleep":
			evt.kind = fallAsleep

		default:
			evt.kind = beginShift
			if _, err := fmt.Sscanf(str, "Guard #%d begins shift", &evt.id); err != nil {
				log.Fatalln("failed to scan event id:", err)
				return nil
			}
		}

		events = append(events, evt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("err scanning input:", err)
		return nil
	}

	return events
}
