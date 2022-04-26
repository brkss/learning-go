package chance

import "math/rand"
import "time"

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano());
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
    var n int;

    n = rand.Intn(19) + 1;
    return (n);
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	var energy float64;

    energy = float64(rand.Intn(12/0.05)) * 0.05;
    return (energy);
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
    var animals []string;

    animals = append(animals, "ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog");
	rand.Shuffle(len(animals), func(i, j int) {
        animals[i], animals[j] = animals[j], animals[i];
    });
	return (animals);
}

