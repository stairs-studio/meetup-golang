package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//Meetup location
var loc = []float64{-27.595377, -48.54805}

//coordinates mock
var coordinates [][]float64

//large distance
var largeDistance float64

func main() {
	getCoordinates()
	start := time.Now()
	l := Coordinate(loc[0], loc[1])
	for _, coordinate := range coordinates {
		c := Coordinate(coordinate[0], coordinate[1])
		d := Distance(l, c)
		if d > largeDistance {
			largeDistance = d
		}
		time.Sleep(time.Microsecond)
	}
	log.Println("time:", time.Since(start))
	log.Println("maior:", largeDistance)
}

func getCoordinates() {
	start := time.Now()
	jsonFile, err := os.Open("../coordinates.json")
	defer jsonFile.Close()
	if err != nil {
		log.Println("file error", err)
		return
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("read error", err)
		return
	}
	err = json.Unmarshal(byteValue, &coordinates)
	if err != nil {
		log.Println("unmarshal error", err)
		return
	}
	log.Println("getCoordinates - time:", time.Since(start))
}
