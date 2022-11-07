package builder

type Director struct {
	builder CakeBuilder
}

func SetDirector(b CakeBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) CookChocalateCake() *Cake {
	d.builder.MixIngredients()
	d.builder.ShakeMixture()
	d.builder.AddMixtureToMold()
	d.builder.CookInOven()
	return d.builder.GetCake()
}
