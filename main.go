package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"kainore.com/palworld/paldex"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var palBreed []paldex.PalBreed

func main() {
	pal := paldex.Load("./data/paldex.json")
	res, err := pal.Find_name("anubis")
	check(err)
	log.Println(res.Name, "found")
	palBreed = paldex.ListFiles()

	router := gin.Default()
	router.GET("/", getAllBreeding)
	router.GET("/byChield/:name", getParrents)
	router.GET("byChieldWithParrent/:chield/:parrent", getMissingParrent)
	router.Run(":8080")
}

func getAllBreeding(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, palBreed)
}

func palBreedContains(name string) (*paldex.PalBreed, error) {
	for _, breed := range palBreed {
		if strings.ToLower(name) == strings.ToLower(breed.Name) {
			return &breed, nil
		}
	}
	return nil, errors.New("Pal not found")
}

func getParrents(c *gin.Context) {
	name := c.Param("name")
	breed, err := palBreedContains(name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, breed)
	return
}

func parrentListContains(list [][]string, value []string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] != nil {
			if list[i][0] == value[0] && list[i][1] == value[1] {
				return true
			}
		}
	}
	return false
}

func findMissingParrent(parrent string, parrents paldex.PalBreed) [][]string {
	var result [][]string
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

func getMissingParrent(c *gin.Context) {
	chield := c.Param("chield")
	parrent := c.Param("parrent")
	parrents, err := palBreedContains(chield)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status":  err.Error(),
			"chield":  chield,
			"parrent": parrent,
		})
		return
	}
	result := findMissingParrent(parrent, *parrents)
	if len(result) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"status":  "Nothing found with this combo",
			"chield":  chield,
			"parrent": parrent,
		})
		return
	}
	c.IndentedJSON(http.StatusNotFound, result)
	return
}
