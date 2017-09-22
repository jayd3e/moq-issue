package main

import (
	"context"
	"github.com/rainhq/moq_test/subpackage"
)

//go:generate moq -out mockpersonstore_test.go . PersonStore

// Person represents a real person.
type Person struct {
	ID      string
	Name    string
	Company string
	Website subpackage.Website
}

// PersonStore provides access to Person objects.
type PersonStore interface {
	Get(ctx context.Context, id string) (*Person, error)
	Create(ctx context.Context, person *Person, confirm bool) error
}

func main() {

}
