package kbo

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL = "https://sports.news.naver.com/kbaseball/schedule/index.nhn"
)

type Parser struct {
	client *http.Client
	url    string
}

// NewParser는 초기화된 Parser의 포인터를 반환합니다.
func NewParser(url string, client *http.Client) *Parser {
	return &Parser{
		client: client,
		url:    url,
	}
}

// Parse는 주어진 시간의 경기 결과를 반환합니다.
func (p *Parser) Parse(t time.Time) ([]Game, error) {
	resp, err := p.request(int(t.Month()), t.Year())
	if err != nil {
		return []Game{}, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return []Game{}, err
	}

	result := p.parse(t.Day(), doc)

	return result, nil
}

func (p *Parser) request(month, year int) (*http.Response, error) {
	req, err := http.NewRequest("GET", p.url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("month", strconv.Itoa(month))
	q.Add("year", strconv.Itoa(year))
	req.URL.RawQuery = q.Encode()

	return p.client.Do(req)
}

func (p *Parser) parse(day int, doc *goquery.Document) []Game {
	games := []Game{}

	doc.Find(".tb_wrap > div[class^=sch_tb]").Each(func(i int, s *goquery.Selection) {
		if Day(s) == day && !NoGame(s) {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				games = append(games, Game{
					Home:     Home(s),
					Away:     Away(s),
					Canceled: Canceled(s),
					Score:    Score(s),
				})
			})
		}
	})

	return games
}

// NoGame은 주어진 selection에서 경기가 취소되었는지 확인합니다.
func NoGame(s *goquery.Selection) bool {
	str, exists := s.Attr("class")
	if !exists {
		return false
	}

	return strings.Contains(str, "nogame")
}

// Day는 주어진 selection이 어떤 날짜를 가지는지 반환합니다.
// 만약 에러가 발생할 경우 -1를 반환합니다.
func Day(s *goquery.Selection) int {
	var (
		day int
		err error
	)
	s.Find(".td_date > strong").Each(func(i int, s *goquery.Selection) {
		day, err = strconv.Atoi(strings.Split(s.Text(), ".")[1])
	})

	if err != nil {
		return -1
	}

	return day
}

// Home은 주어진 selection에서 홈 팀을 반환합니다.
func Home(s *goquery.Selection) Team { return team(s.Find(".team_rgt").Text()) }

// Away는 주어진 selection에서 원정 팀을 반환합니다.
func Away(s *goquery.Selection) Team { return team(s.Find(".team_lft").Text()) }

// Canceled는 주어진 selection에서 경기가 취소되면 true,
// 경기가 취소되지 않으면 false를 반환합니다.
func Canceled(s *goquery.Selection) bool { return s.Find(".td_stadium.cancel").Length() == 1 }

// Score는 주어진 selection에서 경기의 점수를 반환합니다.
// 배열의 첫번째 인덱스는 원정 팀, 두번째 인덱스는 홈 팀의 점수를 가리킵니다.
// 에러가 나면 빈 배열을 반환합니다.
func Score(s *goquery.Selection) [2]int {
	var score [2]int

	if !Canceled(s) {
		strs := strings.Split(s.Find(".td_score").Text(), ":")
		if len(strs) != 2 {
			return [2]int{}
		}

		s1, err := strconv.Atoi(strs[0])
		if err != nil {
			return [2]int{}
		}
		s2, err := strconv.Atoi(strs[1])
		if err != nil {
			return [2]int{}
		}

		score = [2]int{s1, s2}
	}

	return score
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
