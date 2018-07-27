package kbo

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestParserParse(t *testing.T) {
	parser := NewParser(URL, &http.Client{})

	tables := []struct {
		date     time.Time
		expected []Game
	}{
		{
			time.Date(2018, 5, 6, 0, 0, 0, 0, time.UTC),
			[]Game{
				{
					Home:     Lions,
					Away:     Eagles,
					Canceled: true,
					Score:    [2]int{},
				},
				{
					Home:     Wyverns,
					Away:     Giants,
					Canceled: true,
					Score:    [2]int{},
				},
				{
					Home:     Tigers,
					Away:     Dinos,
					Canceled: false,
					Score:    [2]int{3, 11},
				},
				{
					Home:     Twins,
					Away:     Bears,
					Canceled: false,
					Score:    [2]int{13, 5},
				},
				{
					Home:     Wiz,
					Away:     Heroes,
					Canceled: true,
					Score:    [2]int{},
				},
			},
		},
	}

	for _, c := range tables {
		result, _ := parser.Parse(c.date)
		if len(result) != len(c.expected) {
			t.Errorf("length is not same.")
		}

		for i := range result {
			fmt.Println(result[i])
			if result[i] != c.expected[i] {
				t.Errorf("expected:\n%s\ngot:\n%s\nat index %d.", c.expected[i], result[i], i)
			}
		}
	}
}
