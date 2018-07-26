package kbo

const (
	Bears Team = iota + 1
	Dinos
	Eagles
	Giants
	Heroes
	Lions
	Tigers
	Twins
	Wiz
	Wyverns
	Nanum
	Dream
	Unknown
)

type Team int

type Game struct {
	Home     Team
	Away     Team
	Canceled bool
	Score    [2]int
}
