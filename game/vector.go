package game

import "math"

type Vector struct {
	X Float `json:"x"`
	Y Float `json:"y"`
}

func (v Vector) Distance(dest Vector) Vector {
	return Vector {dest.X - v.X, dest.Y - v.Y}
}

func (v Vector) Angle() Float {
	return math.Atan2(v.Y, v.X)
}

func (v *Vector) Add(a Vector) {
	v.X += a.X
	v.Y += a.Y
}

func (v *Vector) Sub(a Vector) {
	v.X -= a.X
	v.Y -= a.Y
}

func (v *Vector) Reset() {
	v.X = 0
	v.Y = 0
}

func (v Vector) SqMagnitude() Float {
	return v.X* v.X + v.Y* v.Y
}

func (v Vector) Magnitude() Float {
	return math.Sqrt(v.SqMagnitude())
}

func Degrees(a Float) Float {
	return a * 180 / math.Pi
}
