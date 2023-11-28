package rotator

import (
	"fmt"
	"math"

	"space-ships/internal/domain"
)

type Rotator struct {
	movable domain.Rotatable
}

func New(movable domain.Rotatable) *Rotator {
	return &Rotator{movable: movable}
}

// Rotate rotates a game object
func (m *Rotator) Rotate() error {
	ang, err := m.movable.GetAngle()
	if err != nil {
		return fmt.Errorf("Rotator: failed to get angle: %w", err)
	}

	vel, err := m.movable.GetAngVelocity()
	if err != nil {
		return fmt.Errorf("Rotator: failed to get angular velocity: %w", err)
	}

	num, err := m.movable.GetDirNumber()
	if err != nil {
		return fmt.Errorf("Rotator: failed to get directions number: %w", err)
	}

	angleStep := float64(360 / num)

	err = m.movable.SetAngle(math.Mod(ang+angleStep*math.Floor(vel/angleStep), 360.0))
	if err != nil {
		return fmt.Errorf("Rotator: failed to update angle: %w", err)
	}

	return nil
}
