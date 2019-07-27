package main

import (
	"KanjiNetInfo/symbol"
	"github.com/jessevdk/go-flags"
)

var cfg struct {
	Char string `long:"char" short:"c" description:"character"`
}

func main() {
	_, err := flags.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	obj := symbol.NewSymbolObj(cfg.Char)
	obj.Find()
	obj.Print()

}
