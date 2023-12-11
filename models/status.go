package models

type Atributte int

type Strength Atributte
type Life Atributte
type Defense Atributte
type Skill Atributte
type Agility Atributte
type Intelligence Atributte

type Status struct {
	Strength     `json:"strength,omitempty" default:"0"`
	Life         `json:"life,omitempty" default:"0"`
	Defense      `json:"defense,omitempty" default:"0"`
	Skill        `json:"skill,omitempty" default:"0"`
	Agility      `json:"agility,omitempty" default:"0"`
	Intelligence `json:"intelligence,omitempty" default:"0"`
}
