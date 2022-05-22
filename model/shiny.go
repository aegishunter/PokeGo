package model

import "fmt"

type shiny interface {
	shine()
}

type ShinyPokemon struct {
	*Pokemon
	shiny
}

func (sp ShinyPokemon) Shine() {
	fmt.Println(sp.GetName() + " is shining!")
}

func NewShinyPokemon(p *Pokemon) ShinyPokemon {
	return ShinyPokemon{Pokemon: p}
}
