package palworld

import (
	"fmt"
	"testing"
)

func TestByChieldName(t *testing.T) {
	name := "Anubis"
	pal := NewPalBreed()
	parrents := pal.ByChieldName(name)
	fmt.Println(len(parrents.Parrents))
	if parrents == nil {
		t.Fatalf("%s doesn't have parrents", name)
	}
}
