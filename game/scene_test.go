package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScene_SurfaceFreeFall(t *testing.T) {
	// This is a quick an dirty soup of numbers that tries to emulate gravity
	// on the surface of Earth, but is far from exact
	const Dim = 2000000000.0 // Indifferent
	const G = 0.00000667259
	const DeltaT = 0.1
	const NIters = 20
	const WeightOfEarth = 5.972 * 1000000000000.0 // In 1000xmillion metric tons (10^12)
	const RadiusOfEarth = 6371000 // in meters
	const WeightOfShip = 1 // Indifferent
	s := NewScene(Dim, Dim, G)
	body := &Body{
		Mass: WeightOfEarth,
	}
	ship := &Ship {Body: Body {
		Pos: Vector{0, RadiusOfEarth + 10},
		//v: Vector{1000, 0},
		Mass: WeightOfShip,
	}}
	s.AddBody(body)
	s.AddPlayer(ship)
	for i := 0; i < NIters; i++ {
		s.Update(DeltaT)
		assert.Zero(t, body.Pos)
		assert.Zero(t, body.V)
		assert.Zero(t, body.f)
		assert.Zero(t, ship.f)
		dist := ship.Pos.Distance(body.Pos).Magnitude() - RadiusOfEarth
		t.Logf("iter %d distance=%f pos=%v", i, dist, ship.Pos)
		if i < 17 {
			// First 1.7 seconds is still in the air
			assert.True(t, dist > 0)
		} else {
			// Then it's below surface (no collisions yet)
			assert.True(t, dist < 0)
		}
	}
	js := string(s.JSon())
	t.Logf("Serialised to %v", js)
}

// 8.69616899342509
func TestScene_Orbit(t *testing.T) {
	// This is a quick an dirty soup of numbers that tries to emulate gravity
	// on the surface of Earth, but is far from exact
	const Dim = 2000000000.0 // Indifferent
	const G = 0.0001
	const DeltaT = 0.1
	//const NIters = 1000
	const WeightOfEarth = 1000000.0 // In 1000xmillion metric tons (10^12)
	const RadiusOfEarth = 1000.0 // in meters
	const WeightOfShip = MinMass // Indifferent
	const Altitude = 1000.0
	s := NewScene(Dim, Dim, G)
	body := &Body{
		Mass: WeightOfEarth,
	}
	ship := &Ship {Body: Body {
		Pos: Vector{0, RadiusOfEarth + Altitude},
		//v: Vector{6221.432103094101, 0.002},
		V:    Vector{X: 8.69616899342509},
		Mass: WeightOfShip,
	}}
	s.AddBody(body)
	s.AddPlayer(ship)
	//for i := 0; i < NIters; i++ {
	wasNegDegs := false
	loops := 0
	for loops < 100 {
		s.Update(DeltaT)
		//assert.Zero(t, body.pos)
		//assert.Zero(t, body.v)
		//assert.Zero(t, body.f)
		//assert.Zero(t, ship.f)
		dist := ship.Pos.Distance(body.Pos).Magnitude() - RadiusOfEarth
		a := Degrees(ship.Pos.Angle())
		if dist < 0 || dist > 2.0*Altitude {
			//return false, a, dist, ship.v.Magnitude() / sp
			return
		}
		negDegs := a < 0
		if wasNegDegs && !negDegs {
			loops ++
			if loops > 0 {
				//return true, a, dist, ship.v.Magnitude() / sp
			}
			t.Logf("loop %d distance=%f deg=%f v=%v", loops, dist, a, ship.V.Magnitude())
		}
		wasNegDegs = negDegs
	}
}
