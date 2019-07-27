package main

import (
	"KanjiNetInfo/symbol"
)

const char string = "質"

func main() {
	obj := symbol.NewSymbolObj(char)
	obj.Find()
	obj.Print()

}
