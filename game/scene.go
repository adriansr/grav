package game

import "math"

// float64 or float32 ?
type Float = float64

type Scene struct {
	bodies   []*Body
	players  []*Ship
	min, max Vector
	g        Float
}

func NewScene(width, height, G Float) *Scene {
	s := &Scene {
		min: Vector{-width/2.0, -height/2.0},
		max: Vector{ width/2.0, height/2.0},
		g:   G,
	}
	return s
}

func (s *Scene) AddBody(body *Body) {
	s.bodies = append(s.bodies, body)
}

func (s *Scene) AddPlayer(player *Ship) {
	s.players = append(s.players, player)
}

func (s *Scene) Update(deltaT Float) {
	// Update gravity b/w large bodies
	for i, n := 0, len(s.bodies); i < n - 1; i++ {
		body := s.bodies[i]
		for j := i + 1; j < n; j++ {
			peer := s.bodies[j]
			f := s.newtonGravity(body, peer)
			body.f.Add(f)
			peer.f.Sub(f)
		}
	}
	// Apply gravity to small bodies
	for _, body := range s.bodies {
		for _, player := range s.players {
			f := s.newtonGravity(&player.Body, body)
			player.f.Add(f)
		}
	}
	// Update speeds
	for _, body := range s.bodies {
		body.ApplyForce(deltaT)
	}
	for _, player := range s.players {
		player.ApplyForce(deltaT)
	}
}

const AlmostZero = 0.000000000000001
func makeZero(v Float) Float {
	if math.Abs(v) < AlmostZero {
		return 0
	}
	return v
}

func (s *Scene) newtonGravity(body, peer *Body) Vector {
	d := body.pos.Distance(peer.pos)
	absdx, absdy := math.Abs(d.x), math.Abs(d.y)
	f := s.g * peer.mass * body.mass / (absdx + absdy)
	sin, cos := math.Sincos(d.Angle())
	return Vector{makeZero(cos * f), makeZero(sin * f)}
}
