package recursion

import (
	"fmt"
)

func PlayMinRefuelStop() {
	stations := [][]int{{25,25}, {50,25}, {75,25}}
	target := 100
	startFuel := 25

	/*
	100
	25
	[[25,25],[50,25],[75,25]]
	 */

	minStops := minRefuelStops(target, startFuel, stations)
	fmt.Println(minStops)
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {

	return minRefuelStopsHelper( target, startFuel, 0, 0, 0, stations)
}

func minRefuelStopsHelper(target, remainingFuel, distance, stops, nextStationId int, stations [][]int) int {
	if remainingFuel < 0 {
		return -1
	}
	if target<=distance+remainingFuel{
		return stops
	}
	//todo
	if nextStationId == len(stations) {
		return -1
	}
	nextStationDistance := stations[nextStationId][0]
	nextStationAvailGas := stations[nextStationId][1]
	remainingFuelAtNext := remainingFuel - nextStationDistance
	//go to next station and refill
	if remainingFuel<nextStationDistance{
		return -1
	}
	s1 := minRefuelStopsHelper(target, remainingFuelAtNext+nextStationAvailGas, distance+nextStationDistance, stops+1, nextStationId+1, stations)
	//go to next station and dont refill
	s2 := minRefuelStopsHelper(target, remainingFuelAtNext, distance+nextStationDistance, stops, nextStationId+1, stations)
	if s1 == -1 {
		return s2
	}
	if s2 == -1 {
		return s1
	}
	return min(s1, s2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
