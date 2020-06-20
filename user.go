package recommender

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// User is a user who recommends something
type User struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Ratings map[string]Rating `json:"ratings"`
}

// NewUser creates and returns a new User
func NewUser(id string, name string) (*User, error) {

	if id == "" {
		id = uuid.NewV4().String()
	}
	return &User{ID: id, Name: name}, nil
}

// String represents a User as a string
func (u User) String() string {
	return fmt.Sprintf("%s", u.Name)
}
