# Go Method

## Method
A mechanism to define a function specialized for an arbitrary type.  
The difference from a normal function is that a receiver is required before the method name.  

## Code
```Go
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

	// Calling the infoMovie method
	movie1.infoMovie()
}
```

## Output Sample
~ $ go build main.go   
~ $ ./main   
Title : Saving Private Ryan, Directer : Steven Spielberg, Year : 1998   
