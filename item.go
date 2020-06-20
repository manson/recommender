package recommender

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Item is and item to recommend
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewItem creates and returns an Item
func NewItem(id string, name string) (*Item, error) {

	if id == "" {
		id = uuid.NewV4().String()
	}
	return &Item{ID: id, Name: name}, nil
}

// String represents an Item as a string
func (i Item) String() string {
	return fmt.Sprintf("%s", i.Name)
}
