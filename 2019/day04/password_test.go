package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePasswordV1(t *testing.T) {
	tests := []struct {
		pw  int
		err error
	}{
		{
			pw:  111111,
			err: nil,
		},
		{
			pw:  223450,
			err: ErrMonotonicDigits,
		},
		{
			pw:  123789,
			err: ErrAdjacentDigits,
		},
		{
			pw:  122345,
			err: nil,
		},
		{
			pw:  111123,
			err: nil,
		},
		{
			pw:  135679,
			err: ErrAdjacentDigits,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			err := ValidatePasswordV1(tt.pw)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestValidatePasswordV2(t *testing.T) {
	tests := []struct {
		pw  int
		err error
	}{
		{
			pw:  111111,
			err: ErrManyAdjacentDigits,
		},
		{
			pw:  223450,
			err: ErrMonotonicDigits,
		},
		{
			pw:  123789,
			err: ErrAdjacentDigits,
		},
		{
			pw:  122345,
			err: nil,
		},
		{
			pw:  111123,
			err: ErrManyAdjacentDigits,
		},
		{
			pw:  135679,
			err: ErrAdjacentDigits,
		},
		{
			pw:  112233,
			err: nil,
		},
		{
			pw:  123444,
			err: ErrManyAdjacentDigits,
		},
		{
			pw:  111122,
			err: nil,
		},
		{
			pw:  111233,
			err: nil,
		},
		{
			pw:  112223,
			err: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			err := ValidatePasswordV2(tt.pw)
			assert.Equal(t, tt.err, err)
		})
	}
}
