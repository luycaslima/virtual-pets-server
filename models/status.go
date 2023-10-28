package models

type Atributte int

type Strength Atributte
type Life Atributte
type Defense Atributte
type Skill Atributte
type Agility Atributte
type Intelligence Atributte

type Status struct {
	Strength     `json:"strength" default:"5"`
	Life         `json:"life" default:"5"`
	Defense      `json:"defense" default:"5"`
	Skill        `json:"skill" default:"5"`
	Agility      `json:"agility" default:"5"`
	Intelligence `json:"intelligence" default:"5"`
}
