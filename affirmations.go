package main

import (
	"fmt"
	"math/rand"
	"time"
)

var affirmations = []string{
	"I am complete as I am, others simply support me",
	"I am grateful for every moment in my life",
	"I have the courage to face all challenges",
	"I attract abundance into my life",
	"I am at peace with myself and the world",
	"I love and accept myself unconditionally",
	"I am constantly growing and evolving",
}

func main() {
	// Seed the random number generator once at the start of the program
	rand.Seed(time.Now().UnixNano())

	// Print a random affirmation
	fmt.Println(RandomAffirmation())
}

// Returns a random affirmation from the list.
func RandomAffirmation() string {
	selectedIndex := rand.Intn(len(affirmations))
	return affirmations[selectedIndex]
}
