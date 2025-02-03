package main

import (
	"testing"
)

func TestAffirmationsHasAtLeastFiveOptions(t *testing.T) {
	affirmationCount := len(affirmations)

	if affirmationCount < 5 {
		t.Errorf("Random affirmation count less than requirement of 5")
	} else {
		t.Logf("There are %d possible affirmations", affirmationCount)
	}
}

func TestRandomAffirmationReturnsAllValidOptions(t *testing.T) {
	affirmationMap := make(map[string]bool)
	for _, a := range affirmations {
		affirmationMap[a] = true
	}
	affirmationCount := len(affirmations)
	foundAffirmations := make(map[string]bool)
	attemptCount := 0
	// Note: This may break if the affirmation count is significantly increased
	// At that time we can implement a more complex validation logic.
	maxAttempts := affirmationCount * 10

	for len(foundAffirmations) < affirmationCount {
		if attemptCount > maxAttempts {
			break
		}
		affirmation := RandomAffirmation()

		if affirmationMap[affirmation] {
			foundAffirmations[affirmation] = true
		} else {
			t.Errorf("Random affirmation was not valid: %s", affirmation)
		}
		attemptCount++
	}

	foundAffirmationCount := len(foundAffirmations)
	if foundAffirmationCount < affirmationCount {
		t.Errorf("Did not return all of the %d possible affirmations in %d attempts", affirmationCount, maxAttempts)
	} else {
		t.Logf("Found %d of %d possible affirmations in %d retrieval attempts", foundAffirmationCount, affirmationCount, attemptCount)
	}
}
