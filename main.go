package main

import (
	"fmt"
	"pokego/model"
)

func main() {

	pikachu := model.NewPokemon("Pikachu", "PKCH", model.NewAttack("Thunder"))
	charmender := model.NewPokemon("Charmender", "CHMR", model.NewAttack("Ember"))
	shinyPikachu := model.NewShinyPokemon(model.NewPokemon("Pikachu", "PKCH", model.NewAttack("Thunder")))
	shinyCharmender := model.NewShinyPokemon(model.NewPokemon("Charmender", "CHMR", model.NewAttack("Ember")))
	hospital := model.NewHospital()
	pikaHospital := model.NewPikaHospital()

	fmt.Println("Heal in Hospital")
	hospital.Heal(pikachu)
	hospital.Heal(shinyPikachu.Pokemon)
	hospital.Heal(charmender)
	hospital.Heal(shinyCharmender.Pokemon)
	fmt.Println("========================")
	fmt.Println()

	fmt.Println("Heal in PikaHospital")
	pikaHospital.Heal(pikachu)
	pikaHospital.Heal(shinyPikachu.Pokemon)
	pikaHospital.Heal(charmender)
	pikaHospital.Heal(shinyCharmender.Pokemon)
	fmt.Println("========================")
	fmt.Println()

	fmt.Println("Before Evolve")
	fmt.Println(pikachu)
	fmt.Println(shinyPikachu.Pokemon)
	fmt.Println(charmender)
	fmt.Println(shinyCharmender.Pokemon)
	fmt.Println("========================")
	fmt.Println()

	fmt.Println("Evolving")
	pikachu.Evolve()
	charmender.Evolve()
	shinyPikachu.Evolve()
	shinyCharmender.Evolve()
	fmt.Println("========================")
	fmt.Println()

	fmt.Println("After evolve")
	fmt.Println(pikachu)
	fmt.Println(shinyPikachu.Pokemon)
	fmt.Println(charmender)
	fmt.Println(shinyCharmender.Pokemon)
	fmt.Println("========================")
	fmt.Println()

	fmt.Println("Shining")
	shinyPikachu.Shine()
	shinyCharmender.Shine()
	fmt.Println("========================")
	fmt.Println()
}
