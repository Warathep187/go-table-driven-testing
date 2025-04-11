package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockChicken struct {
	TakeABathImpl     func() error
	PluckFeathersImpl func() error
}

func (c *MockChicken) TakeABath() error {
	return c.TakeABathImpl()
}

func (c *MockChicken) PluckFeathers() error {
	return c.PluckFeathersImpl()
}

// Take a bath error
// Pluck feathers error
// Success

type EatTestCase struct {
	name             string
	takeABathErr     error
	pluckFeathersErr error
	expectedErr      error
}

func TestHuman_Eat(t *testing.T) {
	tests := []EatTestCase{
		{
			name:             "Take a bath error",
			takeABathErr:     errors.New("failed to take a bath"),
			pluckFeathersErr: nil,
			expectedErr:      errors.New("failed to take a bath"),
		},
		{
			name:             "Pluck feathers error",
			takeABathErr:     nil,
			pluckFeathersErr: errors.New("failed to pluck feathers"),
			expectedErr:      errors.New("failed to pluck feathers"),
		},
		{
			name:             "Success",
			takeABathErr:     nil,
			pluckFeathersErr: nil,
			expectedErr:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockChicken := &MockChicken{}
			mockChicken.TakeABathImpl = func() error {
				return test.takeABathErr
			}
			mockChicken.PluckFeathersImpl = func() error {
				return test.pluckFeathersErr
			}

			human := Human{}
			err := human.Eat(mockChicken)

			assert.Equal(t, err, test.expectedErr)
		})
	}
}
