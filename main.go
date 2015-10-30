package main

import (
	"log"

	"github.com/jeffail/gabs"
)

func main() {
	jsonParsed, _ := gabs.ParseJSON([]byte(`{
    "smaller":100000,
    "bigger":1000000
  }`))

	log.Println("=>", jsonParsed.String())
}
