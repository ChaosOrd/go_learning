package main

import (
	"fmt"
)

type Planets []string

func (p Planets) Terraform() {

	for idx, planet := range p {
		p[idx] = "New" + planet
	}
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	Planets(planets).Terraform()
	fmt.Println(planets)
}