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
	Home     Team   `json:"home"`
	Away     Team   `json:"away"`
	Canceled bool   `json:"canceled"`
	Score    [2]int `json:"score"`
}
