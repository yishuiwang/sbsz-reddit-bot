package basic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	Servers = "https://api.bgm.tv"
)

type SubjectType int

const (
	Books SubjectType = iota + 1
	Animation
	Music
	Games
	_
	ThreeDimensional
)

type Search struct {
	Results int    `json:"results"`
	List    []List `json:"list"`
}
type Character struct {
	BirthMon  int         `json:"birth_mon"`
	Gender    string      `json:"gender"`
	BirthDay  int         `json:"birth_day"`
	BirthYear interface{} `json:"birth_year"`
	BloodType interface{} `json:"blood_type"`
	Images    Images      `json:"images"`
	Summary   string      `json:"summary"`
	Name      string      `json:"name"`
	Infobox   []Infobox   `json:"infobox"`
	Stat      Stat        `json:"stat"`
	ID        int         `json:"id"`
	Locked    bool        `json:"locked"`
	Type      int         `json:"type"`
	Nsfw      bool        `json:"nsfw"`
}
type Anime struct {
	Date          string     `json:"date"`
	Platform      string     `json:"platform"`
	Images        Images     `json:"images"`
	Summary       string     `json:"summary"`
	Name          string     `json:"name"`
	NameCn        string     `json:"name_cn"`
	Tags          []Tags     `json:"tags"`
	Infobox       []Infobox  `json:"infobox"`
	Rating        Rating     `json:"rating"`
	TotalEpisodes int        `json:"total_episodes"`
	Collection    Collection `json:"collection"`
	ID            int        `json:"id"`
	Eps           int        `json:"eps"`
	Volumes       int        `json:"volumes"`
	Locked        bool       `json:"locked"`
	Nsfw          bool       `json:"nsfw"`
	Type          int        `json:"type"`
}
type List struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	Type       int    `json:"type"`
	Name       string `json:"name"`
	NameCn     string `json:"name_cn"`
	Summary    string `json:"summary"`
	AirDate    string `json:"air_date"`
	AirWeekday int    `json:"air_weekday"`
	Images     Images `json:"images"`
}
type Images struct {
	Small  string `json:"small"`
	Grid   string `json:"grid"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Common string `json:"common"`
}
type Tags struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
type Infobox struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Rating struct {
	Rank  int     `json:"rank"`
	Total int     `json:"total"`
	Score float64 `json:"score"`
}
type Collection struct {
	OnHold  int `json:"on_hold"`
	Dropped int `json:"dropped"`
	Wish    int `json:"wish"`
	Collect int `json:"collect"`
	Doing   int `json:"doing"`
}
type Stat struct {
	Comments int `json:"comments"`
	Collects int `json:"collects"`
}

func SubjectInfo(id string) (Anime, error) {
	path := "/v0/subjects/" + id
	res, err := http.Get(Servers + path)
	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var a Anime
	json.Unmarshal(data, &a)
	return a, err
}

func CharacterInfo(id string) (Character, error) {
	path := "/v0/characters/" + id
	res, err := http.Get(Servers + path)
	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var c Character
	json.Unmarshal(data, &c)
	return c, err
}

func SearchKeyWord(keyword string, category int) (Search, error) {

	path := "/search/subject/" + keyword
	switch category {
	case 1:
		path += fmt.Sprintf("?type=%d", Books)
	case 2:
		path += fmt.Sprintf("?type=%d", Animation)
	case 3:
		path += fmt.Sprintf("?type=%d", Music)
	case 4:
		path += fmt.Sprintf("?type=%d", Games)
	case 6:
		path += fmt.Sprintf("?type=%d", ThreeDimensional)
	default:
		path += fmt.Sprintf("?type=%d", Animation)
	}

	res, err := http.Get(Servers + path)

	data, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var s Search
	json.Unmarshal(data, &s)
	return s, err
}
