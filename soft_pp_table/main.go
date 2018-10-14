package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	file  = kingpin.Arg("file", "File to read hex values from").Required().String()
	cards = kingpin.Arg("card", "Card # to write to").Required().String()
)

func main() {
	kingpin.Parse()
	b, err := ioutil.ReadFile(*file)
	ppTable := strings.TrimSpace(string(b))
	fmt.Printf("%s\n", ppTable)
	b, err = hex.DecodeString(ppTable)
	if err != nil {
		log.Fatalf("Failed to decode string: %v", err)
	}

	cardsStrList := strings.Split(*cards, ",")
	for _, cardStr := range cardsStrList {
		card, err := strconv.Atoi(cardStr)
		if err != nil {
			log.Fatalf("Bad card #: %v: %v", cardStr, err)
		}
		ioutil.WriteFile("pp", b, 0664)
		ppTablePath := fmt.Sprintf("/sys/class/drm/card%d/device/pp_table", card)
		if err := ioutil.WriteFile(ppTablePath, b, 0666); err != nil {
			log.Fatalf("Failed to write pp_table file: %v", err)
		}
	}
}
