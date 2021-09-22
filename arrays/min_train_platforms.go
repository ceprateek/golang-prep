package arrays

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func PlayFindPlatforms() {
	arrivals := []string{"9:00", "9:40", "9:50", "11:00", "15:00", "18:00"}
	departures := []string{"9:10", "12:00", "11:20", "11:30", "19:00", "20:00"}
	p := findMinPlatforms(arrivals, departures)
	fmt.Println(p)
}

func getTime(trainTime string) time.Time {
	t := time.Now()
	year, month, day := t.Date()
	timeMidnight := time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	splits := strings.Split(trainTime, ":")
	h, _ := strconv.Atoi(splits[0])
	m, _ := strconv.Atoi(splits[1])
	tt := timeMidnight.Add(time.Hour*time.Duration(h) + time.Minute*time.Duration(m))
	return tt
}

func findMinPlatforms(arrivals, departures []string) (platforms int) {
	arrivalTimes := make([]time.Time, len(arrivals))
	departureTimes := make([]time.Time, len(departures))
	for i, val := range arrivals {
		arrivalTimes[i] = getTime(val)
	}
	for i, val := range departures {
		departureTimes[i] = getTime(val)
	}
	sort.Slice(arrivalTimes, func(i, j int) bool {
		return arrivalTimes[i].Before(arrivalTimes[j])
	})
	sort.Slice(departureTimes, func(i, j int) bool {
		return departureTimes[i].Before(departureTimes[j])
	})
	i := 0
	j := 0
	platforms = 0
	maxPlatforms := 0
	for i < len(arrivalTimes) && j < len(departureTimes) {
		if arrivalTimes[i].Before(departureTimes[j]) || arrivalTimes[i].Equal(departureTimes[j]){
			platforms++
			i++
			maxPlatforms = int(math.Max(float64(platforms), float64(maxPlatforms)))
		}else {
			platforms--
			j++
		}
	}

	return maxPlatforms
}
