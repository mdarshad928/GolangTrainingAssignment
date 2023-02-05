/*
Q.10 Create a structure named country_cap with members country & capital city (Population).

type country_cap struct {
	country string
	population int
}

Declare an array of 10 elements and initialize it with data.
	a) Sort the array by country name in ascending order and print the elements.
	b) Sort the array by population in descending order and print the elements.

	Hint: use Sort function of the package “sort” & implement sort.Interface for each sorting criteria.
*/
package main

import (
	"fmt"
	"sort"
)

type country_cap struct {
	country    string
	population int
}

// ByPopulation implements sort.Interface for []country_cap based on
// the population field.
type ByPopulation []country_cap

func (a ByPopulation) Len() int           { return len(a) }
func (a ByPopulation) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPopulation) Less(i, j int) bool { return a[i].population < a[j].population }

// ByCountry implements sort.Interface for []country_cap based on
// the country field.
type ByCountry []country_cap

func (a ByCountry) Len() int           { return len(a) }
func (a ByCountry) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCountry) Less(i, j int) bool { return a[i].country < a[j].country }

func main() {
	country_ca := [10]country_cap{{"India", 1352642280}, {"Brazil", 215730676}, {"China", 1411750000}, {"Indonesia", 1411750000},
		{"USA", 334233854}, {"Japan", 1252642280}, {"Bangladesh", 165158616}, {"Mexico", 128665641}, {"Russia", 146424729}, {"Nigeria", 218541000}}
	country_cap := country_ca[:]
	sort.Sort(ByCountry(country_cap))
	fmt.Println(country_cap)
}
