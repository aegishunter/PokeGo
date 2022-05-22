package model

type Attack struct {
	name string
}

func (a Attack) GetName() string {
	return a.name
}

func NewAttack(name string) Attack {
	return Attack{name: name}
}
