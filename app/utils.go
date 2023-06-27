package app

import (
	"sort"
)

// Implementing Generics would avoid repetition displayed here.
// TODO: Implement Generics for these functions.

func getHighestFloat(requests []float64) float64 {
	// Initialize the highest float.
	highest := requests[0]

	// Loop through the requests and find the highest one.
	for _, request := range requests {
		if request > highest {
			highest = request
		}
	}

	return highest
}

func getHighestInt(requests []int) int {
	// Initialize the highest int.
	highest := requests[0]

	// Loop through the requests and find the highest one.
	for _, request := range requests {
		if request > highest {
			highest = request
		}
	}

	return highest
}

func getLowestFloat(requests []float64) float64 {
	// Initialize the lowest float.
	lowest := requests[0]

	// Loop through the requests and find the lowest one.
	for _, request := range requests {
		if request < lowest {
			lowest = request
		}
	}

	return lowest
}

func getLowestInt(requests []int) int {
	// Initialize the lowest int.
	lowest := requests[0]

	// Loop through the requests and find the lowest one.
	for _, request := range requests {
		if request < lowest {
			lowest = request
		}
	}

	return lowest
}

func getMeanOfRequests(requests []float64) float64 {
	// Initialize the sum of the requests.
	sum := 0.0

	// Loop through the requests and add them to the sum.
	for _, request := range requests {
		sum += request
	}

	mean := sum / float64(len(requests))
	return mean
}

func getMedianRequest(requests []float64) float64 {
	// Sort the requests.
	sort.Float64s(requests)

	// Initialize the median.
	median := 0.0

	if len(requests)%2 == 0 {
		median = (requests[len(requests)/2-1] + requests[len(requests)/2]) / 2
	} else {
		median = requests[len(requests)/2]
	}

	return median
}

/*
The percentage of requests that succeeded
Any error codes returned that weren't a success
*/
