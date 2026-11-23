package main

import "fmt"

func DistanceMeasure(seats []int) int {
	var count, maxDist int
	seenPerson := false

	for _, seat := range seats {
		if seat == 0 {
			count++
		} else {
			if !seenPerson {
				maxDist = count
				seenPerson = true
			} else {
				dist := (count + 1) / 2
				if dist > maxDist {
					maxDist = dist
				}
			}
			count = 0
		}
	}

	if count > maxDist {
		maxDist = count
	}

	return maxDist
}

func main() {
	fmt.Println(DistanceMeasure([]int{1, 0, 0, 0, 1, 0, 1}))
	fmt.Println(DistanceMeasure([]int{1, 0, 0, 0}))
	fmt.Println(DistanceMeasure([]int{0, 0, 0, 1}))
	fmt.Println(DistanceMeasure([]int{0, 1}))
	fmt.Println(DistanceMeasure([]int{1, 0, 0, 0, 0, 1}))
}
