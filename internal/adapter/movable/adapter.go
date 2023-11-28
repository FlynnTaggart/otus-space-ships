package movable

import (
	"errors"
	"fmt"

	"space-ships/internal/domain"
)

const (
	propertyPosition = "Position"
	propertyVelocity = "Velocity"
)

var (
	errFailCastPosition = errors.New("MovableAdapter: failed to cast object position to vector")
	errFailCastVelocity = errors.New("MovableAdapter: failed to cast object velocity to vector")
)

type Adapter struct {
	obj domain.UObject
}

func New(obj domain.UObject) *Adapter {
	return &Adapter{obj: obj}
}

func (a *Adapter) GetPosition() (domain.Vector, error) {
	rawPos, err := a.obj.GetProperty(propertyPosition)
	if err != nil {
		return domain.Vector{}, fmt.Errorf("MovableAdapter: failed to get position: %w", err)
	}

	pos, ok := rawPos.(domain.Vector)
	if !ok {
		return domain.Vector{}, errFailCastPosition
	}

	return pos, nil
}

func (a *Adapter) GetVelocity() (domain.Vector, error) {
	rawVel, err := a.obj.GetProperty(propertyVelocity)
	if err != nil {
		return domain.Vector{}, fmt.Errorf("MovableAdapter: failed to get velocity: %w", err)
	}

	vel, ok := rawVel.(domain.Vector)
	if !ok {
		return domain.Vector{}, errFailCastVelocity
	}

	return vel, nil
}

func (a *Adapter) SetPosition(newPos domain.Vector) error {
	err := a.obj.SetProperty(propertyPosition, newPos)
	if err != nil {
		return fmt.Errorf("MovableAdapter: failed to set position: %w", err)
	}

	return nil
}
