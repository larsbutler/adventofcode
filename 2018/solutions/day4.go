package solutions

import (
	"fmt"
	"sort"
	"strings"
)

type GuardId int
type Date struct {
	month int
	day int
}
/*
type GuardShift struct {
	guardId GuardId
	month int
	day int
	minutesAsleep []int
}
*/

// Map hierarchy: guardID -> date -> minutes ([]int of 60 0 or 1 values)
type Minutes [60]int
type GuardDailyMinutes map[Date]*Minutes
type GuardShifts map[GuardId]GuardDailyMinutes

func MinuteMax(m Minutes) (int, int) {
	var maxI int = 0
	var max int = m[0]
	for i := 1; i < len(m); i++ {
		if m[i] > max {
			max = m[i]
			maxI = i
		}
	}
	return maxI, max
}

func AddMinutes(a Minutes, b Minutes) Minutes {
	var result Minutes = Minutes{}
	for i := 0; i < len(result); i++ {
		result[i] = a[i] + b[i]
	}
	return result
}

func SumMinutes(m Minutes) int {
	var sum int = 0
	for _, x := range m {
		sum += x
	}
	return sum
}

func maxIndexMinutes(m Minutes) int {
	var max int = 0
	var maxIndex int = -1
	for i, x := range m {
		if x >= max {
			max = x
			maxIndex = i
		}
	}
	return maxIndex
}

func printShifts(gs GuardShifts) {
	for gid, gdm := range gs {
		for date, minutes := range gdm {
			fmt.Printf("%d-%d\t#%d\t%v\n", date.month, date.day, gid, *minutes)
		}
	}
}

func getGuardIdWithMostMinutesAsleep(gs GuardShifts) GuardId {
	var mostMinutes map[GuardId]int = make(map[GuardId]int)

	for gid, gdm := range gs {
		for _, minutes := range gdm {
			var subtotal = SumMinutes(*minutes)
			if _, exists := mostMinutes[gid]; !exists {
				mostMinutes[gid] = subtotal
			} else {
				mostMinutes[gid] += subtotal
			}
		}
	}
	fmt.Println(mostMinutes)

	var target GuardId = -1
	var max int = 0
	for gid, total := range mostMinutes {
		if total > max {
			max = total
			target = gid
		}
	}

	return target
}

func getMinuteGuardSleepsMost(gs GuardShifts, gid GuardId) int {
	// return value is an index for the minute, i.e.: 0-59
	var gdm GuardDailyMinutes = gs[gid]
	var overlapping Minutes = Minutes{}

	// Overlay all of the minutes together and keep a running tally of
	// total minutes slept at a given time of day.
	// For example:
	//   [0, 1, 1, 0]
	// + [1, 0, 1, 0]
	// + [0, 0, 1, 0]
	// = [1, 1, 3, 0]
	// In this case, minute index 2 is the one with the most frequent
	// naps.

	for _, minutes := range gdm {
		for i := 0; i < len(overlapping); i++ {
			overlapping[i] += minutes[i]
		}
	}

	fmt.Printf("Overlapping: %v\n", overlapping)
	return maxIndexMinutes(overlapping)
}

func getGuardWhoSleptMostDuringGivenMinute(gs GuardShifts) (GuardId, int) {
	var targetGuard GuardId
	var targetMinute int
	var maxSoFar int = 0

	for gid, gdm := range gs {
		var m Minutes = Minutes{}
		for _, minutes := range gdm {
			m = AddMinutes(m, *minutes)
		}
		// Get the minute where each guard slept the most:
		var maxI int
		var max int
		maxI, max = MinuteMax(m)
		if max > maxSoFar {
			targetGuard = gid
			targetMinute = maxI
			maxSoFar = max
		}
	}
	fmt.Printf("%v, %v\n", targetGuard, targetMinute)
	return targetGuard, targetMinute
}

func storeShift(gs GuardShifts, gid GuardId, date Date, timeAsleep int, timeAwake int) {
	// Check if guard id already exists
	if gdm, exists := gs[gid]; !exists {
		// Guard doesn't yet have a record; make one!
		gdm = make(map[Date]*Minutes)
		gs[gid] = gdm
	}

	var gdm GuardDailyMinutes = gs[gid]
	// See if the date yet exists:
	if entry, exists := gdm[date]; !exists {
		// Date doesn't yet exist for this guard.
		// Create the array of 60:
		entry = &Minutes{}
		gdm[date] = entry
	}

	var entry *Minutes = gdm[date]
	// The minute the guard falls asleep and the minute before
	// they wake up count to the total.
	// On the minute the guard wakes up, they are considered "awake",
	// and so this minute doesn't count toward the total.
	for i := timeAsleep; i < timeAwake; i++ {
		(*entry)[i] = 1
	}

}


func parseDate(s string) (month int, day int) {
	var parts []string = strings.Split(s, "-")
	month = AsInt(parts[1])
	day = AsInt(parts[2])
	return month, day
}

func parseTime(s string) (hour int, minute int) {
	var parts []string = strings.Split(s, ":")
	hour = AsInt(parts[0])
	minute = AsInt(parts[1])
	return hour, minute
}

func getGuardShifts(lines []string) GuardShifts {

	var guardShifts GuardShifts = make(map[GuardId]GuardDailyMinutes)
	var guardId GuardId = 0
	var timeAsleep int = 0  // time the guards fall asleep
	for _, line := range lines {
		var parts []string = strings.Split(line, " ")
		// part one is the date
		var date string = strings.Trim(parts[0], "[")
		var month, day int = parseDate(date)

		// part two is the time
		var time string = strings.Trim(parts[1], "]")
		var _, minute int = parseTime(time)

		fmt.Println(line)
		if parts[2] == "Guard" && parts[4] == "begins" && parts[5] == "shift" {
			// TODO: close off the last guard shift?
			// track shift hours for this guard
			guardId = GuardId(AsInt(strings.Trim(parts[3], "#")))
		} else if parts[2] == "falls" && parts[3] == "asleep" {
			// TODO: guard falls asleep
			timeAsleep = minute
		} else if parts[2] == "wakes" && parts[3] == "up" {
			// TODO: guard wakes up
			// `minute` right now is the time awake
			var timeAwake int = minute
			var date Date = Date{month: month, day: day}
			storeShift(guardShifts, guardId, date, timeAsleep, timeAwake)
		}
	}
	return guardShifts
}

func Day4Part2(input string) string {
	var lines []string = SplitLines(input)
	sort.Slice(
		lines,
		func(i, j int) bool {
			return lines[i] < lines[j]
		},
	)
	var guardShifts GuardShifts = getGuardShifts(lines)

	// find which guard spent the most time asleep during any given
	// minute
	// which guard and which minute?
	var gid GuardId
	var minute int
	gid, minute = getGuardWhoSleptMostDuringGivenMinute(guardShifts)

	return AsStr(int(gid) * minute)
}

func Day4Part1(input string) string {
	// 1. Sort the input by time
	var lines []string = SplitLines(input)
	sort.Slice(
		lines,
		func(i, j int) bool {
			return lines[i] < lines[j]
		},
	)
	var guardShifts GuardShifts = getGuardShifts(lines)
	printShifts(guardShifts)
	var gwm GuardId = getGuardIdWithMostMinutesAsleep(guardShifts)
	var theMinute int = getMinuteGuardSleepsMost(guardShifts, gwm)
	fmt.Printf("Guard with most minutes asleep: %d\n", gwm)
	fmt.Printf("Guard sleeps the most at minute %d\n", theMinute)

	// Solution for Part 1 is:
	// guardId * minute the guard sleeps the most
	return AsStr(int(gwm) * theMinute)
}
