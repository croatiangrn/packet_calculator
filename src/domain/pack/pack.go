package pack

import "fmt"

// Pack represents a package with an ID and item size
type Pack struct {
	ID       int
	ItemSize int
}

// Validate checks if the Pack has a valid item size
func (p *Pack) Validate() error {
	if p.ItemSize <= 0 {
		return fmt.Errorf("item size must be greater than zero")
	}

	return nil
}
