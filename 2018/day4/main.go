package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type eventType int

const (
	GuardBeginsDuty eventType = iota
	FallsAsleep
	Awakes
)

var (
	eventNames = [...]string{
		"Guard Begins Duty",
		"Falls Asleep",
		"Wakes Up",
	}
)

func (et eventType) String() string {
	return eventNames[et]
}

type event struct {
	at      time.Time
	typ     eventType
	guardId int
}

func parseEvent(s string) (event, error) {
	var year, month, day, hour, minute, guardId int
	var msg string
	fmt.Sscanf(s, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)
	msg = s[19:]
	t := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
	if strings.HasPrefix(msg, "Guard #") {
		fmt.Sscanf(msg, "Guard #%d begins shift", &guardId)
		return event{t, GuardBeginsDuty, guardId}, nil
	} else if msg == "falls asleep" {
		return event{t, FallsAsleep, 0}, nil
	} else if msg == "wakes up" {
		return event{t, Awakes, 0}, nil
	} else {
		return event{}, fmt.Errorf("error creating event from: %s, unknown event", s)
	}
}

func main() {
	events := make([]event, 0, 934)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		event, err := parseEvent(scanner.Text())
		if err != nil {
			fmt.Printf("ERROR reading input: %s\n", err.Error())
			os.Exit(-1)
		}
		events = append(events, event)
	}
	sort.SliceStable(events, func(i, j int) bool {
		before := events[i].at.Before(events[j].at)
		if before {
			return true
		}
		if events[i].at.Equal(events[j].at) {
			return events[i].typ < events[j].typ
		}
		return false
	})
	var guardId int
	var timeAsleep time.Time

	guardSleepHeatMap := make(map[int]*[60]int)
	for _, event := range events {
		if event.typ == GuardBeginsDuty {
			guardId = event.guardId
		}
		guardHeatMap, ok := guardSleepHeatMap[guardId]
		if !ok {
			guardHeatMap = new([60]int)
			guardSleepHeatMap[guardId] = guardHeatMap
		}
		fmt.Printf("%s %s %d\n", event.at.String(), event.typ.String(), guardId)
		if event.typ == FallsAsleep {
			timeAsleep = event.at
		} else if event.typ == Awakes {
			fmt.Printf("Processing Awake event %v, %v\n", timeAsleep, event.at)
			for m := timeAsleep.Minute(); m < event.at.Minute(); m += 1 {
				guardHeatMap[m] = guardHeatMap[m] + 1
			}
		}
	}
	highestMinutes := 0
	guardSelected := -1
	minuteSelected := -1
	for guardId, guardHeatMap := range guardSleepHeatMap {
		for idx, c := range guardHeatMap {
			if c > highestMinutes {
				highestMinutes = c
				minuteSelected = idx
				guardSelected = guardId
			}
		}
	}
	fmt.Printf("Guard #%d, Minute: %d, answer %d\n", guardSelected, minuteSelected, guardSelected*minuteSelected)
}
