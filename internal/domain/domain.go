package domain

//go:generate mockery --name UObject

// UObject Common game object
type UObject interface {
	GetProperty(key string) (any, error)
	SetProperty(key string, value any) error
}

type Movable interface {
	GetPosition() (Vector, error)
	GetVelocity() (Vector, error)
	SetPosition(Vector) error
}

type Rotatable interface {
	GetAngle() (float64, error)
	GetAngVelocity() (float64, error)
	SetAngle(float64) error
	GetDirNumber() (int64, error) // GetDirNumber Number of direction that could be set to object
}
