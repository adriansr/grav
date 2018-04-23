package game

const MinMass = 0.0005

type Body struct {
	Pos    Vector `json:"pos"`
	V      Vector `json:"v"`
	f      Vector // TODO: Maybe keep f outside?
	Mass   Float `json:"mass"`
	Radius Float `json:"radius"`
}

func (b *Body) ApplyForce(deltaT Float) {
	if b.Mass < MinMass {
		return
	}
	baseV := b.V
	ax := b.f.X / b.Mass
	ay := b.f.Y / b.Mass
	b.V.X += ax * deltaT
	b.V.Y += ay * deltaT
	b.Pos.X += (b.V.X + baseV.X) * deltaT / 2.0
	b.Pos.Y += (b.V.Y + baseV.Y) * deltaT / 2.0
	// Reset force after applying
	b.f.Reset()
}
