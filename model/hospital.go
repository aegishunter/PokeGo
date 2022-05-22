package model

import "fmt"

type Hospital struct {
}

type Heal interface {
	Heal()
}

func (h Hospital) Heal(pokemon *Pokemon) {
	fmt.Println(pokemon.GetName() + " is healed!")
}

func NewHospital() *Hospital {
	return &Hospital{}
}
