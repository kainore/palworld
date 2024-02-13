package paldex

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type PalBreed struct {
	Name     string     `json:"chield"`
	Parrents [][]string `json:"parrents"`
}

func ListFiles() []PalBreed {
	r, err := os.ReadDir("./data/pals/")
	if err != nil {
		log.Fatal(err)
	}

	var breedings []PalBreed
	dex := Load("./data/paldex.json")
	for i := 0; i < len(r); i++ {

		f := r[i]

		chield := f.Name()
		p, err := dex.Find_name(chield)
		if err != nil {
			if strings.Contains(chield, "(special)") {
			} else if strings.Contains(chield, "ice kingpaca") {
				p.Name = "Kingpaca Cryst"
			} else if strings.Contains(chield, "ice reptyro") {
				p.Name = "Reptyro Cryst"
			} else {
				log.Fatal("Couldn't find pal ", chield, "\n", err)
			}
		}
		data, err := os.ReadFile(fmt.Sprintf("./data/pals/%s", chield))
		chield = p.Name
		if err != nil {
			log.Fatal(err)
		}

		var parrents [][]string

		parr := strings.Split(string(data), "\n")

		for _, names := range parr {
			temp := chieldToParrents(names)
			if len(temp) == 2 {
				parrents = append(parrents, temp)
			}
		}
		pb := PalBreed{
			Name:     chield,
			Parrents: parrents,
		}
		breedings = append(breedings, pb)
	}
	return breedings
}

func chieldToParrents(names string) []string {
	arr := trim(strings.Split(names, "+"))
	return arr
}

func trim(arr []string) []string {
	for i, name := range arr {
		arr[i] = strings.TrimSpace(name)
	}
	return arr
}
