package main

import (
	"fmt"
	"strconv"
)

// Planet Interface implemented by all planets
type Planet interface {
	name() string
	mass() int
}

// Mercury Planet
type Mercury struct {
	// TODO donot keep fields  just return something from methods - Done
}

func (planet Mercury) name() string {
	return "Mercury"
}
func (planet Mercury) mass() int {
	return 1
}
func (planet Mercury) String() string {
	return planet.name() + " " + strconv.Itoa(planet.mass())
}

// Venus Planet
type Venus struct {
}

func (planet Venus) name() string {
	return "Venus"
}
func (planet Venus) mass() int {
	return 2
}
func (planet Venus) String() string {
	return planet.name() + " " + strconv.Itoa(planet.mass())
}

// Earth Planet
type Earth struct {
}

func (planet Earth) name() string {
	return "Earth"
}
func (planet Earth) mass() int {
	return 3
}
func (planet Earth) String() string {
	return planet.name() + " " + strconv.Itoa(planet.mass())
}

// Mars Planet
type Mars struct {
}

func (planet Mars) name() string {
	return "Mars"
}
func (planet Mars) mass() int {
	return 4
}
func (planet Mars) String() string {
	return planet.name() + " " + strconv.Itoa(planet.mass())
}

func main() {

	// TODO slice of planets
	// make a function input arg is planet
	planets := []Planet{
		Mercury{},
		Venus{},
		Earth{},
		Mars{},
	}

	for _, planet := range planets {
		fmt.Println(planetInfo(planet))
	}
}

func planetInfo(planet Planet) string {
	return fmt.Sprintf("Name:%s, Mass:%d", planet.name(), planet.mass())
}
