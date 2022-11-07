//go:generate mockgen -source iBuilder.go -destination mock/builder_mock.go -package mock;
package builder

type CakeBuilder interface {
	MixIngredients()
	ShakeMixture()
	AddMixtureToMold() error
	CookInOven()
	GetCake() *Cake
}

func GetBuilder(builderType string) CakeBuilder {
	if builderType == "ChocolateLarge" {
		return LargeChocolateCakeBuilder()
	}
	if builderType == "ChocolateSmall" {
		return SmallChocolateCakeBuilder()
	}

	return nil
}
