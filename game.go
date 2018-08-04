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

func team(s string) Team {
	switch s {
	case "KIA":
		return Tigers
	case "두산":
		return Bears
	case "LG":
		return Twins
	case "SK":
		return Wyverns
	case "롯데":
		return Giants
	case "한화":
		return Eagles
	case "NC":
		return Dinos
	case "KT":
		return Wiz
	case "넥센":
		return Heroes
	case "삼성":
		return Lions
	case "나눔":
		return Nanum
	case "드림":
		return Dream
	default:
		return Unknown
	}
}
