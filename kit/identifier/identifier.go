package identifier

import "github.com/google/uuid"

// New generate an unique identifier.
func New() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// Parse from a primitive value returns identifier if is valid.
func Parse(value string) (string, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}

	return id.String(), err
}