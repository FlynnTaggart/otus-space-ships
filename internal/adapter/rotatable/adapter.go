package rotatable

import (
	"errors"
	"fmt"

	"space-ships/internal/domain"
)

const (
	propertyAngle       = "Angle"
	propertyAngVelocity = "AngVelocity"
	propertyDirNumber   = "DirNumber"
)

var (
	errFailCastAngle       = errors.New("RotatableAdapter: failed to cast object angle to float64")
	errFailCastAngVelocity = errors.New("RotatableAdapter: failed to cast object angular velocity to float64")
	errFailCastDirNum      = errors.New("RotatableAdapter: failed to cast object directions number to int64")
)

// Adapter over UObject for rotation
type Adapter struct {
	obj domain.UObject
}

func New(obj domain.UObject) *Adapter {
	return &Adapter{obj: obj}
}

func (a *Adapter) GetAngle() (float64, error) {
	rawAng, err := a.obj.GetProperty(propertyAngle)
	if err != nil {
		return 0, fmt.Errorf("RotatableAdapter: failed to get angle: %w", err)
	}

	ang, ok := rawAng.(float64)
	if !ok {
		return 0, errFailCastAngle
	}

	return ang, nil
}

func (a *Adapter) GetAngVelocity() (float64, error) {
	rawVel, err := a.obj.GetProperty(propertyAngVelocity)
	if err != nil {
		return 0, fmt.Errorf("RotatableAdapter: failed to get angular velocity: %w", err)
	}

	vel, ok := rawVel.(float64)
	if !ok {
		return 0, errFailCastAngVelocity
	}

	return vel, nil
}

func (a *Adapter) SetAngle(newAngle float64) error {
	err := a.obj.SetProperty(propertyAngle, newAngle)
	if err != nil {
		return fmt.Errorf("RotatableAdapter: failed to set angle: %w", err)
	}

	return nil
}

// GetDirNumber Number of direction that could be set to object
func (a *Adapter) GetDirNumber() (int64, error) {
	rawNum, err := a.obj.GetProperty(propertyDirNumber)
	if err != nil {
		return 0, fmt.Errorf("RotatableAdapter: failed to get directons number: %w", err)
	}

	num, ok := rawNum.(int64)
	if !ok {
		return 0, errFailCastDirNum
	}

	return num, nil
}
