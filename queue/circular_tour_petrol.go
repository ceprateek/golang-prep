package queue

import "fmt"

/*
For example, let there be 4 petrol pumps with amount of petrol and distance
to next petrol pump value pairs as {4, 6}, {6, 5}, {7, 3} and {4, 5}.
The first point from where the truck can make a circular tour is 2nd petrol pump.
Output should be “start = 1” (index of 2nd petrol pump).
*/

type petrolPump struct {
	petrol, distance int
}

func printTour(pumps []petrolPump) int {
	start := 0
	end := 1
	petrolRemaining := pumps[start].petrol - pumps[start].distance
	for end != start || petrolRemaining < 0 {
		for petrolRemaining < 0 && start != end {
			petrolRemaining = petrolRemaining - (pumps[start].petrol - pumps[start].distance)
			start = (start + 1) % len(pumps)
			if start == 0 {
				return -1
			}
		}
		petrolRemaining += pumps[end].petrol - pumps[end].distance
		end = (end + 1) % len(pumps)
	}
	return start
}

func PlayCircularPetrolPump() {
	pumps := make([]petrolPump, 0, 4)
	p0 := petrolPump{
		petrol:   4,
		distance: 6,
	}
	pumps = append(pumps, p0)
	p1 := petrolPump{
		petrol:   6,
		distance: 5,
	}
	pumps = append(pumps, p1)
	p2 := petrolPump{
		petrol:   7,
		distance: 3,
	}
	pumps = append(pumps, p2)
	p3 := petrolPump{
		petrol:   4,
		distance: 5,
	}
	pumps = append(pumps, p3)
	startPump := printTour(pumps)
	fmt.Printf("startPump: %d", startPump)
}
