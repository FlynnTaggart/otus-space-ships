package mover

import (
	"fmt"

	"space-ships/internal/domain"
)

type Mover struct {
	movable domain.Movable
}

func New(movable domain.Movable) *Mover {
	return &Mover{movable: movable}
}

// Move moves a game object
func (m *Mover) Move() error {
	pos, err := m.movable.GetPosition()
	if err != nil {
		return fmt.Errorf("Mover: failed to get position: %w", err)
	}

	vel, err := m.movable.GetVelocity()
	if err != nil {
		return fmt.Errorf("Mover: failed to get velocity: %w", err)
	}

	err = m.movable.SetPosition(domain.AddVectors(pos, vel))
	if err != nil {
		return fmt.Errorf("Mover: failed to update position: %w", err)
	}

	return nil
}
