package main

import (
	"fmt"
)

// Struct for movie titles
type Movie struct {
	Directer string
	Title    string
	Year     int
}

// Methods to output movie information
func (m *Movie) infoMovie() {
	// Output via receiver variables
	fmt.Printf("Title : %s, Directer : %s, Year : %d \n", m.Title, m.Directer, m.Year)
}

// Main function
func main() {

	// Stores movie information
	movie1 := &Movie{
		Directer: "Steven Spielberg",
		Title: "Saving Private Ryan",
		Year: 1998,
	}
	movie2 := &Movie{
		Directer: "Akira Kurosawa",
		Title: "Seven Samurai",
		Year: 1954,
	}
	movie3 := &Movie{
		Directer: "Roman Polanski",
		Title: "The Pianist",
		Year: 2002,
	}

	// Calling the infoMovie method
	movie1.infoMovie()
	movie2.infoMovie()
	movie3.infoMovie()
}

// =================================
//           Output Sample
// =================================
// ~ $ go build main.go 
// ~ $ ./main 
// Title : Saving Private Ryan, Directer : Steven Spielberg, Year : 1998 
// Title : Seven Samurai, Directer : Akira Kurosawa, Year : 1954 
// Title : The Pianist, Directer : Roman Polanski, Year : 2002 
