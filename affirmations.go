package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Add new affirmations here
    selfWorth := "I am complete as I am, others simply support me"
    gratitude := "I am grateful for every moment in my life"
    courage := "I have the courage to face all challenges"
    abundance := "I attract abundance into my life"
    peace := "I am at peace with myself and the world"
    selfLove := "I love and accept myself unconditionally"
    growth := "I am constantly growing and evolving"

    // All affirmations must be stored in a slice
    affirmations := []string{selfWorth, gratitude, courage, abundance, peace, selfLove, growth}

    //Seed the random number generator to ensure randomness
    rand.Seed(time.Now().UnixNano())

    //Select and print a random affirmation
    selectedIndex := rand.Intn(len(affirmations))
    fmt.Println(affirmations[selectedIndex])
}