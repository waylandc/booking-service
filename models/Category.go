package models

// Category is used to group our  services for searchability
type Category struct {
	ID   int
	Name string `json:"name"`
}
