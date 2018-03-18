package models

// Organization is a service provider
type Organization struct {
	ID         int
	Name       string     `json:"org_name"`
	Website    string     `json:"website"`
	Phone      string     `json:"phone"`
	City       string     `json:"city"`
	Categories []Category `json:"categories"`
}
