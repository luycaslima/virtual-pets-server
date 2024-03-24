package models

type Atributte uint

type Strength Atributte
type Life Atributte
type Defense Atributte
type Skill Atributte
type Agility Atributte
type Intelligence Atributte

type Status struct {
	Strength     `json:"strength,omitempty" `
	Life         `json:"life,omitempty" `
	Defense      `json:"defense,omitempty" `
	Skill        `json:"skill,omitempty" `
	Agility      `json:"agility,omitempty" `
	Intelligence `json:"intelligence,omitempty" `
}
