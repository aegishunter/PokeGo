package model

import "fmt"

type Pokemon struct {
	name     string
	attacks  []Attack
	poketype string
}

func (p Pokemon) GetName() string {
	return p.name
}

func (p *Pokemon) Evolve() {
	if p.name == "Pikachu" {
		p.name = "Raichu"
		p.attacks = append(p.attacks, NewAttack("Mega Thunder"))
		fmt.Println("Pikachu has evolved into Raichu")
	} else {
		fmt.Println(p.name + " can't be evolve right now :(")
	}
}

func NewPokemon(name string, poketype string, attacks ...Attack) *Pokemon {
	return &Pokemon{
		name:     name,
		attacks:  attacks,
		poketype: poketype,
	}
}
