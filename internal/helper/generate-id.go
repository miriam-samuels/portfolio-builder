package helper

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func GenerateUniqueId(length int) string {
	// Define the character set of digits (0-9)
	characters := "0123456789"

	// Create a new source for random number generation
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Create a byte slice to store the generated ID
	id := make([]byte, length)

	// Generate the ID by randomly selecting digits from the character set
	for i := range id {
		id[i] = characters[r.Intn(len(characters))]
	}

	// Convert the byte slice to a string
	return string(id)
}
