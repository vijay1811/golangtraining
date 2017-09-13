package main

import (
	"fmt"
	"testing"
)

func nameCheck(expected, actual string) error {
	if expected != actual {
		return fmt.Errorf("expected: %v, actual: %v", expected, actual)
	}
	return nil
}

func massCheck(expected, actual int) error {
	if expected != actual {
		return fmt.Errorf("expected: %v, actual: %v", expected, actual)
	}
	return nil
}
func printCheck(expected, actual string) error {
	if expected != actual {
		return fmt.Errorf("expected: %v, actual: %v", expected, actual)
	}
	return nil
}

func Test_nameMercury(t *testing.T) {
	merc := Mercury{}
	if err := nameCheck("Mercury", merc.name()); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_massMercury(t *testing.T) {
	mercury := Mercury{}
	if err := massCheck(1, mercury.mass()); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_stringMercury(t *testing.T) {
	mercury := Mercury{}
	actual := fmt.Sprintf("Name: %v, Mass: %v", mercury.name(), mercury.mass())
	if err := printCheck("Name: Mercury, Mass: 1", actual); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_nameVenus(t *testing.T) {
	venus := Venus{}
	if err := nameCheck("Venus", venus.name()); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_massVenus(t *testing.T) {
	venus := Venus{}
	if err := massCheck(2, venus.mass()); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_stringVenus(t *testing.T) {
	venus := Venus{}
	actual := fmt.Sprintf("Name: %v, Mass: %v", venus.name(), venus.mass())
	if err := printCheck("Name: Venus, Mass: 2", actual); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_nameEarth(t *testing.T) {
	earth := Earth{}
	if err := nameCheck("Earth", earth.name()); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_massEarth(t *testing.T) {
	earth := Earth{}
	if err := massCheck(3, earth.mass()); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_StringEarth(t *testing.T) {
	earth := Earth{}
	actual := fmt.Sprintf("Name: %v, Mass: %v", earth.name(), earth.mass())
	if err := printCheck("Name: Earth, Mass: 3", actual); err != nil {
		t.Fatalf("%v", err)
	}
}
func Test_nameMars(t *testing.T) {
	mars := Mars{}
	if err := nameCheck("Mars", mars.name()); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_massMars(t *testing.T) {
	mars := Mars{}
	if err := massCheck(4, mars.mass()); err != nil {
		t.Fatalf("%v", err)
	}
}

func Test_StringMars(t *testing.T) {
	mars := Mars{}
	actual := fmt.Sprintf("Name: %v, Mass: %v", mars.name(), mars.mass())
	if err := printCheck("Name: Mars, Mass: 4", actual); err != nil {
		t.Fatalf("%v", err)
	}
}
