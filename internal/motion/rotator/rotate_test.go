package rotator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"space-ships/internal/adapter/rotatable"
	"space-ships/internal/domain/mocks"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(obj *mocks.UObject)
		err     string
	}{
		{
			name: "success",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Angle").Return(90.0, nil)
				obj.EXPECT().GetProperty("AngVelocity").Return(90.0, nil)
				obj.EXPECT().GetProperty("DirNumber").Return(int64(8), nil)
				obj.EXPECT().SetProperty("Angle", 180.0).Return(nil)
			},
		},
		{
			name: "fail get angle",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Angle").Return(270.0,
					errors.New("no such key"))
			},
			err: "get angle",
		},
		{
			name: "fail get angular velocity",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Angle").Return(270.0, nil)
				obj.EXPECT().GetProperty("AngVelocity").Return(130.0,
					errors.New("no such key"))
			},
			err: "get angular velocity",
		},
		{
			name: "fail get directions number",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Angle").Return(270.0, nil)
				obj.EXPECT().GetProperty("AngVelocity").Return(130.0, nil)
				obj.EXPECT().GetProperty("DirNumber").Return(int64(8),
					errors.New("no such key"))
			},
			err: "get directions number",
		},
		{
			name: "fail update angle",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Angle").Return(270.0, nil)
				obj.EXPECT().GetProperty("AngVelocity").Return(130.0, nil)
				obj.EXPECT().GetProperty("DirNumber").Return(int64(8), nil)
				obj.EXPECT().SetProperty("Angle", 0.0).Return(errors.New("no such key"))
			},
			err: "update angle",
		},
		{
			name: "success over limit",
			prepare: func(obj *mocks.UObject) {
				obj.EXPECT().GetProperty("Angle").Return(270.0, nil)
				obj.EXPECT().GetProperty("AngVelocity").Return(160.0, nil)
				obj.EXPECT().GetProperty("DirNumber").Return(int64(8), nil)
				obj.EXPECT().SetProperty("Angle", 45.0).Return(nil)
			},
		},
	}
	for _, test := range tests {
		obj := mocks.NewUObject(t)

		test.prepare(obj)

		adapter := rotatable.New(obj)
		rotator := New(adapter)

		err := rotator.Rotate()

		if len(test.err) != 0 {
			assert.ErrorContains(t, err, test.err)
			continue
		}
		assert.NoError(t, err)
	}
}
