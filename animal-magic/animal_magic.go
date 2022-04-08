package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	randomed := rand.Intn(20) + 1
	return randomed
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	randomed := rand.Float64() * 12
	return randomed
}

/*
ShuffleAnimals returns a slice with all eight animal strings in random order

The game features eight different animals:

- ant
- beaver
- cat
- dog
- elephant
- fox
- giraffe
- hedgehog */
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}
