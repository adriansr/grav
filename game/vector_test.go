package game

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector_Distance(t *testing.T) {
	for idx, testCase := range []struct {
		source, dest, expected Vector
	} {
		{},
		{ Vector{0, 0}, Vector{1, 0}, Vector{1, 0}},
		{ Vector{0, 0}, Vector{0, 2}, Vector{0, 2}},
		{ Vector{10, 0}, Vector{0, 0}, Vector{-10, 0}},
		{ Vector{0, 20}, Vector{0, 0}, Vector{0, -20}},
		{ Vector{10, 10}, Vector{-5, -40}, Vector{-15, -50}},
		{ Vector{-3, 7}, Vector{100, -25}, Vector{103, -32}},
		{ Vector{-7.5, -1.125}, Vector{-4.75, 25.78125}, Vector{2.75, 26.90625}},
	} {
		errMsg := fmt.Sprintf("Test case %d: %+v", idx+1, testCase)
		assert.Equal(t, testCase.expected, testCase.source.Distance(testCase.dest), errMsg)
	}
}

func TestVector_Angle(t *testing.T) {
	const EXACT = -1
	const DELTA = 0.0000005
	for idx, testCase := range []struct {
		delta Float
		vector Vector
		expected Float
	} {
		{EXACT,Vector{1, 0}, 0},
		{EXACT,Vector{100, 0}, 0},
		{EXACT,Vector{0, 1}, math.Pi / 2},
		{EXACT,Vector{0, 50}, math.Pi / 2},
		{EXACT, Vector{0.5, 0.5}, math.Pi / 4},
		{EXACT, Vector{-7.5, 0}, math.Pi},
		{EXACT, Vector{0, -123.125}, -math.Pi / 2},
		{EXACT, Vector{13, -13}, -math.Pi / 4},
		{EXACT, Vector{-100, -100}, -3 * math.Pi / 4},
		{EXACT, Vector{-8, 8}, 3 * math.Pi / 4},
		{DELTA, Vector{-1, 0.5}, 2.6779449},
	} {
		errMsg := fmt.Sprintf("Test case %d: %+v", idx+1, testCase)
		result := testCase.vector.Angle()
		if testCase.delta < 0 {
			assert.Equal(t, testCase.expected, result, errMsg)
		} else {
			assert.InDelta(t, testCase.expected, result, testCase.delta, errMsg)
		}
	}
}

func TestVector_AddSub(t *testing.T) {
	for idx, testCase := range []struct {
		a, b, c Vector
	} {
		{},
		{ Vector{0, 0}, Vector{0, 0}, Vector{0, 0}},
		{ Vector{1, 2}, Vector{-3, -4}, Vector{-2, -2}},
	} {
		errMsg := fmt.Sprintf("Test case %d: %+v", idx+1, testCase)
		r := testCase
		r.a.Add(r.b)
		assert.Equal(t, r.c, r.a, errMsg)
		r = testCase
		r.b.Add(r.a)
		assert.Equal(t, r.c, r.b, errMsg)
		r = testCase
		r.a.Add(r.b)
		assert.Equal(t, r.c, r.a, errMsg)
		r = testCase
		r.c.Sub(r.a)
		assert.Equal(t, r.b, r.c, errMsg)
		r = testCase
		r.c.Sub(r.b)
		assert.Equal(t, r.a, r.c, errMsg)
	}
}
