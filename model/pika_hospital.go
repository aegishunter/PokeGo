package model

import "fmt"

type PikaHospital struct {
}

func (ph PikaHospital) Heal(pokemon *Pokemon) {
	if pokemon.poketype != "PKCH" {
		fmt.Println("You can't heal this pokemon here!")
	} else {
		fmt.Println(pokemon.GetName() + " is healed!")
	}
}

func NewPikaHospital() *PikaHospital {
	return &PikaHospital{}
}
