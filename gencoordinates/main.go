package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"
)

//Meetup location
var loc = []float64{-27.595377, -48.54805}

//Max range between location and one generated coordinate (km)
var maxRange float64 = 100

//Range latitude
var latRange = []float64{-30, -25}

//Range longitude
var lonRange = []float64{-50, -45}

//Num of generated coordinates
var num = 1000000

var coordinates = [][]float64{}

func main() {
	start := time.Now()
	for len(coordinates) < num {
		lat := random(latRange[0], latRange[1])
		lon := random(lonRange[0], lonRange[1])
		d := Distance(&loc[0], &loc[1], &lat, &lon)
		if d < maxRange {
			coordinates = append(coordinates, []float64{
				lat, lon,
			})
		}
	}
	bytes, err := json.MarshalIndent(coordinates, "", "	")
	if err != nil {
		log.Println("marshal error", err)
		return
	}
	err = ioutil.WriteFile("../coordinates.json", bytes, 0644)
	if err != nil {
		log.Println("writefile error", err)
		return
	}
	log.Println("time:", time.Since(start))
}

func random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return ((max - min) * rand.Float64()) + min
}

//Distance - distance between two points
func Distance(lat1, lon1, lat2, lon2 *float64) float64 {
	return 2 * math.Atan2(math.Sqrt(math.Pow(math.Sin((*lat2*math.Pi/180-*lat1*math.Pi/180)/2), 2)+math.Cos(*lat1*math.Pi/180)*math.Cos(*lat2*math.Pi/180)*math.Pow(math.Sin((*lon2*math.Pi/180-*lon1*math.Pi/180)/2), 2)), math.Sqrt(1-(math.Pow(math.Sin((*lat2*math.Pi/180-*lat1*math.Pi/180)/2), 2)+
		math.Cos(*lat1*math.Pi/180)*math.Cos(*lat2*math.Pi/180)*
			math.Pow(math.Sin((*lon2*math.Pi/180-*lon1*math.Pi/180)/2), 2)))) * 6378.16
}
