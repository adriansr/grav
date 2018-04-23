package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody_ApplyForce(t *testing.T) {
	for idx, testCase := range []struct {
		mass Float
		pos Vector
		v Vector
		f Vector
		deltaT Float
		expectedPos Vector
	} {
		{deltaT:1},
		{1000, Vector{100, -25}, Vector{33, 0}, Vector{-66000, 0}, 0.5, Vector{100, -25}},
		{1000, Vector{100, -25}, Vector{33, 0}, Vector{-264000, 52000}, 0.5, Vector{1, 1}},
	}{
		errMsg := fmt.Sprintf("Test case %d: %+v", idx+1, testCase)
		body := Body{
			mass: testCase.mass,
			pos: testCase.pos,
			v: testCase.v,
			f: testCase.f,
		}
		body.ApplyForce(testCase.deltaT)
		assert.Equal(t, testCase.expectedPos, body.pos, errMsg)
		// Force should be always zero after being applied
		assert.Zero(t, body.f)
	}
}
