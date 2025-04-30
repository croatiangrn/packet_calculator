package pack

import "fmt"

type Pack struct {
	ID       int
	ItemSize int
}

func (p *Pack) Validate() error {
	if p.ItemSize <= 0 {
		return fmt.Errorf("item size must be greater than zero")
	}

	return nil
}
