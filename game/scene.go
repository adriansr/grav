package game

import (
	"encoding/json"
	"math"
)

// float64 or float32 ?
type Float = float64

type Scene struct {
	Bodies   []*Body `json:"bodies"`
	Players  []*Ship `json:"players"`
	Min      Vector `json:"universe_min"`
	Max Vector `json:"universe_max"`
	G        Float
}

type InternalError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func NewScene(width, height, G Float) *Scene {
	s := &Scene {
		Min: Vector{-width/2.0, -height/2.0},
		Max: Vector{ width/2.0, height/2.0},
		G:   G,
	}
	return s
}

func (s *Scene) AddBody(body *Body) {
	s.Bodies = append(s.Bodies, body)
}

func (s *Scene) AddPlayer(player *Ship) {
	s.Players = append(s.Players, player)
}

func (s *Scene) Update(deltaT Float) {
	// Update gravity b/w large bodies
	for i, n := 0, len(s.Bodies); i < n - 1; i++ {
		body := s.Bodies[i]
		for j := i + 1; j < n; j++ {
			peer := s.Bodies[j]
			f := s.newtonGravity(body, peer)
			body.f.Add(f)
			peer.f.Sub(f)
		}
	}
	// Apply gravity to small bodies
	for _, body := range s.Bodies {
		for _, player := range s.Players {
			f := s.newtonGravity(&player.Body, body)
			player.f.Add(f)
		}
	}
	// Update speeds
	for _, body := range s.Bodies {
		body.ApplyForce(deltaT)
	}
	for _, player := range s.Players {
		player.ApplyForce(deltaT)
	}
}

func (s *Scene) JSon() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		data, err = json.Marshal(InternalError{})
		if err != nil {
			data = []byte("\"double fault\"")
		}
	}
	return data
}

const AlmostZero = 0.000000000000001
func makeZero(v Float) Float {
	if math.Abs(v) < AlmostZero {
		return 0
	}
	return v
}

func (s *Scene) newtonGravity(body, peer *Body) Vector {
	d := body.Pos.Distance(peer.Pos)
	absdx, absdy := math.Abs(d.X), math.Abs(d.Y)
	f := s.G * peer.Mass * body.Mass / (absdx + absdy)
	sin, cos := math.Sincos(d.Angle())
	return Vector{makeZero(cos * f), makeZero(sin * f)}
}
