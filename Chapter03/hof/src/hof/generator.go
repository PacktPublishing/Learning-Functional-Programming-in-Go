package hof

import (
	"sync"
	"log"
)


func carGenerator(iterator func(int) int, lower int, upper int) func() (int, bool) {
	return func() (int, bool) {
		lower = iterator(lower)
		return lower, lower > upper
	}
}

func iterator(i int) int {
	i += 1
	return i
}


func (cars Collection) GenerateCars(start, limit int) Collection {

	// Create an unbuffered channel to receive match carChannel to display.
	carChannel := make(chan *IndexedCar)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	numCarsToGenerate := start + limit - 1
	generatedCars := Collection{}

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(numCarsToGenerate)

	next := carGenerator(iterator, start -1, numCarsToGenerate)

	carIndex, done := next()
	for !done {
		// Launch the goroutine to perform the search.
		go func(carIndex int) {
			thisCar, err := GetThisCar(carIndex)
			//log.Printf("GetThisCar(%v): %v\n", carIndex, thisCar)
			if err != nil {
				panic(err)
			}
			carChannel <- thisCar
			generatedCars = append(generatedCars, thisCar.Car)
			waitGroup.Done()
		}(carIndex)

		carIndex, done = next()
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel to signal to the getCars
		// function that we can exit the program.
		close(carChannel)
	}()


	// Start displaying carChannel as they are available and
	// return after the final result is displayed.
	printCars(carChannel, start, limit)
	return generatedCars
}



// getCars writes results to the console window as they
// are received by the individual goroutines.
func printCars(indexedCars chan *IndexedCar, start, limit int) {
	log.Printf("\nGenerated Cars (%d to %d)\n%s\n", start, start + limit, DASHES)
	// The channel blocks until a car is written to the channel.
	// Once the channel is closed the for loop terminates.
	var cars Collection
	for car := range indexedCars {
		log.Printf("car: %s\n", car.Car)
		cars = append(cars, car.Car)
	}
}
