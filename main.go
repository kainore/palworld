package main

import (
	"log"
	"time"

	"kainore.com/palworld/paldex"
)

type Pal struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	paldex := paldex.Load("./data/paldex.json")
	res, err := paldex.Find_name("anubis")
	check(err)
	log.Println(res.Name, "found")
}

func measure(f func([]int, int) int, arr []int, num int) time.Duration {
	start := time.Now()
	f(arr, num)
	end := time.Now()
	return end.Sub(start)
}

func measure_sort(f func([]int), arr []int) time.Duration {
	start := time.Now()
	f(arr)
	return time.Now().Sub(start)
}

func measure_sort_r(f func([]int) []int, arr []int) ([]int, time.Duration) {
	start := time.Now()
	arr2 := f(arr)
	return arr2, time.Now().Sub(start)
}
