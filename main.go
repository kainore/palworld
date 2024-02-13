package palworld

import (
	"log"
	"strings"

	"github.com/kainore/palworld/paldex"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type PalBreed struct {
	PalBreed []paldex.PalBreed `json:"pal_breed"`
}

func NewPalBreed() *PalBreed {
	palBreed := paldex.ListFiles()
	return &PalBreed{
		PalBreed: palBreed,
	}
}

func (pal *PalBreed) GetAll() *[]paldex.PalBreed {
	return &pal.PalBreed
}

func (pal *PalBreed) ByChieldName(name string) *paldex.PalBreed {
	for _, breed := range pal.PalBreed {
		if strings.ToLower(name) == strings.ToLower(breed.Name) {
			return &breed
		}
	}
	return nil
}

func (pal *PalBreed) FindByChieldAndParrent(chield, parrent string) [][]string {
	var result [][]string
	parrents := pal.ByChieldName(chield)
	for i := 0; i < len(parrents.Parrents); i++ {
		for j := 0; j < len(parrents.Parrents[i]); j++ {
			if strings.ToLower(parrent) == strings.ToLower(parrents.Parrents[i][j]) {
				if !parrentListContains(result, parrents.Parrents[i]) {
					result = append(result, parrents.Parrents[i])
				}
			}
		}
	}

	return result
}

func parrentListContains(list [][]string, value []string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] != nil {
			if list[i][0] == value[0] && list[i][1] == value[1] {
				return true
			} else if list[i][1] == value[0] && list[i][0] == value[0] {
				return true
			}
		}
	}
	return false
}
