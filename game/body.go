package game

const MinMass = 0.0005

type Body struct {
	pos    Vector
	v      Vector
	f      Vector // TODO: Maybe keep f outside?
	mass   Float
	radius Float
}

func (b *Body) ApplyForce(deltaT Float) {
	if b.mass < MinMass {
		return
	}
	baseV := b.v
	ax := b.f.x / b.mass
	ay := b.f.y / b.mass
	b.v.x += ax * deltaT
	b.v.y += ay * deltaT
	b.pos.x += (b.v.x + baseV.x) * deltaT / 2.0
	b.pos.y += (b.v.y + baseV.y) * deltaT / 2.0
	// Reset force after applying
	b.f.Reset()
}
