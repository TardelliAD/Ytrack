package main

import (
	"fmt"
)

type eleve struct {
	nom string
	age int
}

func (s *eleve) Init(nom string, age int) {
	s.nom = nom
	s.age = age
}
func (s eleve) Affichage() {
	fmt.Println("nom:", s.nom, "\nage:", s.age)
}
func main() {
	var s1 eleve
	s1.Init("andrea", 17)
	s1.Affichage()
}
