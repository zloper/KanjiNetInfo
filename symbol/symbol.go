package symbol

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

type JsonConf struct {
	UrlGoogle string `json:"urlGoogle"`
	Wiki      string `json:"wiki"`
	Youtube   string `json:"youtube"`
}

type Symbol struct {
	Kanji      string
	UrlKanji   string
	GoogleUrl  string
	WikiURL    string
	YoutubeUrl string
	GoogleRaw  string
}

type SymbolMethods interface {
	Find()
	Print()
}

func NewSymbolObj(character string) SymbolMethods {
	urlCharacter := &url.URL{Path: character}
	return &Symbol{Kanji: character, UrlKanji: urlCharacter.String()}
}

func (s *Symbol) Find() {
	var cfg JsonConf
	f, _ := os.Open("./config.json")
	defer f.Close()

	bJson, _ := ioutil.ReadAll(f)
	err := json.Unmarshal(bJson, &cfg)
	if err != nil {
		panic(err)
	}

	s.GoogleUrl = strings.Replace(cfg.UrlGoogle, "{paste}", s.UrlKanji+"+kanji", -1)
	s.WikiURL = cfg.Wiki + s.UrlKanji
	s.YoutubeUrl = cfg.Youtube + s.UrlKanji
	s.GoogleRaw = strings.Replace(cfg.UrlGoogle, "{paste}", s.UrlKanji, -1)
}

func (s *Symbol) Print() {
	tmp := s.GoogleUrl
	if tmp != "" {
		fmt.Println(tmp)
	}

	tmp = s.WikiURL
	if tmp != "" {
		fmt.Println(tmp)
	}

	tmp = s.YoutubeUrl
	if tmp != "" {
		fmt.Println(tmp)
	}

	tmp = s.GoogleRaw
	if tmp != "" {
		fmt.Println(tmp)
	}
}
