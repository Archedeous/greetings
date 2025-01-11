package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Seed a new random generator and log its seed.
	seed := time.Now().UnixNano()
	fmt.Println("Seed:", seed) // Debugging output
	rnd := rand.New(rand.NewSource(seed))

	// Get the random index and log it.
	index := rnd.Intn(len(formats))
	fmt.Println("Selected index:", index) // Debugging output

	return formats[index]
}
