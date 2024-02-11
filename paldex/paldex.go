package paldex

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type PalDex struct {
	Pals []Pal
}

func Load(path string) *PalDex {
	var pals []Pal
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteData, &pals)

	return &PalDex{Pals: pals}
}

func (p *PalDex) Find_name(name string) (Pal, error) {
	var pal Pal
	found := false
	name = strings.Map(unicode.ToLower, name)
	for i := 0; i < len(p.Pals); i++ {
		if name == strings.Map(unicode.ToLower, p.Pals[i].Name) {
			found = true
			pal = p.Pals[i]
			break
		}
	}
	if !found {
		return pal, errors.New("No pal found")
	}
	return pal, nil
}
