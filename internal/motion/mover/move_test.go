package mover

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"space-ships/internal/adapter/movable"
	"space-ships/internal/domain"
	"space-ships/internal/domain/mocks"
)

func TestMove(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(obj *mocks.UObject)
		err     string
	}{
		{
			name: "success",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Position").Return(domain.Vector{X: 12, Y: 5}, nil)
				obj.EXPECT().GetProperty("Velocity").Return(domain.Vector{X: -7, Y: 3}, nil)
				obj.EXPECT().SetProperty("Position", domain.Vector{X: 5, Y: 8}).Return(nil)
			},
		},
		{
			name: "fail get position",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Position").Return(domain.Vector{X: 12, Y: 5},
					errors.New("no such key"))
			},
			err: "get position",
		},
		{
			name: "fail get velocity",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Position").Return(domain.Vector{X: 12, Y: 5}, nil)
				obj.EXPECT().GetProperty("Velocity").Return(domain.Vector{X: -7, Y: 3},
					errors.New("no such key"))
			},
			err: "get velocity",
		},
		{
			name: "fail update position",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Position").Return(domain.Vector{X: 12, Y: 5}, nil)
				obj.EXPECT().GetProperty("Velocity").Return(domain.Vector{X: -7, Y: 3}, nil)
				obj.EXPECT().SetProperty("Position", domain.Vector{X: 5, Y: 8}).
					Return(errors.New("no such key"))
			},
			err: "update position",
		},
	}
	for _, test := range tests {
		obj := mocks.NewUObject(t)

		test.prepare(obj)

		adapter := movable.New(obj)
		mover := New(adapter)

		err := mover.Move()

		if len(test.err) != 0 {
			assert.ErrorContains(t, err, test.err)
			continue
		}
		assert.NoError(t, err)
	}
}
