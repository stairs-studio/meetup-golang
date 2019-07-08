package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

//Meetup location
var loc = []float64{-27.595377, -48.54805}

//coordinates mock
var coordinates [][]float64

//largest distance
var largestDistance float64

//Main function
func main() {
	//Get coordinates from json file
	getCoordinates()

	//Start timer
	start := time.Now()

	//Transform Meetup location in a calculable object
	l := Coordinate(loc[0], loc[1])

	//Make a channel of jobs
	jobs := make(chan []float64, 1000)

	//Create a wait group sync
	wg := sync.WaitGroup{}

	//Initialize workers with a loop
	for w := 0; w < 100; w++ {

		//Anonymous function in a routine
		go func() {

			//Listening jobs channel
			for j := range jobs {

				//Transform coordinate in a calculable object
				c := Coordinate(j[0], j[1])

				//Calculate distance between locations
				d := Distance(l, c)

				//Verify largest distance
				if d > largestDistance {
					largestDistance = d
				}

				//Sleep a microsecond to do the work more expensive
				time.Sleep(time.Microsecond)

				//Decrease a counter from waiter
				wg.Done()
			}
		}()
	}

	//Range coordinates
	for _, coordinate := range coordinates {

		//Increase a counter from waiter
		wg.Add(1)

		//Put data in channel
		jobs <- coordinate
	}

	//Wait the counter go to 0
	wg.Wait()

	//Print the runtime
	log.Println("time:", time.Since(start))

	//Print the largest distance
	log.Println("maior:", largestDistance)
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
