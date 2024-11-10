package main

import (
	"fmt"
)

func main() {
	// Example population data for each year.
	// You can modify this list with actual data.
	populationData := map[int]int{
		2000: 1000000,
		2001: 1020000,
		2002: 1045000,
		2003: 1070000,
		2004: 1100000,
		2005: 1135000,
	}

	// Calculate the total population over the years
	var totalPopulation int
	for _, population := range populationData {
		totalPopulation += population
	}

	// Display total population
	fmt.Printf("Total Population: %d\n\n", totalPopulation)

	// Calculate and display the percentage for each year
	fmt.Printf("Year | Population | Percentage of Total Population\n")
	fmt.Println("-----------------------------------------------------")
	for year, population := range populationData {
		percentage := (float64(population) / float64(totalPopulation)) * 100
		fmt.Printf("%d   | %d         | %.2f%%\n", year, population, percentage)
	}
}
