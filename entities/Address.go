package entities

// Address is an entity
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
