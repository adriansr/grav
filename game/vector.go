package game

import "math"

type Vector struct {
	x, y float64
}

func (v Vector) Distance(dest Vector) Vector {
	return Vector {dest.x - v.x, dest.y - v.y}
}

func (v Vector) Angle() Float {
	return math.Atan2(v.y, v.x)
}

func (v *Vector) Add(a Vector) {
	v.x += a.x
	v.y += a.y
}

func (v *Vector) Sub(a Vector) {
	v.x -= a.x
	v.y -= a.y
}

func (v *Vector) Reset() {
	v.x = 0
	v.y = 0
}

func (v Vector) SqMagnitude() Float {
	return v.x * v.x + v.y * v.y
}

func (v Vector) Magnitude() Float {
	return math.Sqrt(v.SqMagnitude())
}

func Degrees(a Float) Float {
	return a * 180 / math.Pi
}
