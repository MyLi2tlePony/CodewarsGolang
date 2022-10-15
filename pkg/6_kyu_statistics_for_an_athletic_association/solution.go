package statistics_for_an_athletic_association

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func solution(str string) string {
	if str == "" {
		return ""
	}

	times := strings.Split(str, ",")

	results := make([]time.Time, len(times))

	for index, currentTime := range times {
		currentTime = strings.Trim(currentTime, " ")
		result := strings.Split(currentTime, "|")

		hour, _ := strconv.Atoi(result[0])
		minute, _ := strconv.Atoi(result[1])
		second, _ := strconv.Atoi(result[2])

		results[index] = time.Date(0, 0, 0, hour, minute, second, 0, time.UTC)
	}

	resultsRange := calculateRange(results)
	strRange := fmt.Sprintf("Range: %02d|%02d|%02d", resultsRange.Hour(), resultsRange.Minute(), resultsRange.Second())

	resultAverage := calculateAverage(results)
	strAverage := fmt.Sprintf("Average: %02d|%02d|%02d", resultAverage.Hour(), resultAverage.Minute(), resultAverage.Second())

	resultMedian := calculateMedian(results)
	strMedian := fmt.Sprintf("Median: %02d|%02d|%02d", resultMedian.Hour(), resultMedian.Minute(), resultMedian.Second())

	return fmt.Sprintf("%s %s %s", strRange, strAverage, strMedian)
}

func calculateMedian(times []time.Time) time.Time {
	sort.Slice(times, func(i int, j int) bool {
		return times[i].Before(times[j])
	})

	l := len(times)

	if l%2 == 0 {
		return times[l/2-1].Add(times[l/2].Sub(times[l/2-1]) / 2)
	}

	return times[l/2]
}

func calculateAverage(times []time.Time) time.Time {
	var seconds int

	for _, t := range times {
		seconds += t.Second() + t.Minute()*60 + t.Hour()*60*60
	}

	l := len(times)

	return time.Date(0, 0, 0, 0, 0, seconds/l, 0, time.UTC)
}

func calculateRange(times []time.Time) time.Time {
	maxTime := times[0]
	minTime := times[0]

	for _, currentTime := range times {
		if currentTime.After(maxTime) {
			maxTime = currentTime
		}
		if currentTime.Before(minTime) {
			minTime = currentTime
		}
	}

	resultRange := int(maxTime.Sub(minTime).Seconds())
	return time.Date(0, 0, 0, 0, 0, resultRange, 0, time.UTC)
}
