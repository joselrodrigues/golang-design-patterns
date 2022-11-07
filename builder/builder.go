package builder

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Size string

var (
	SMALL  Size = "S"
	MEDIUM Size = "M"
	LARGE  Size = "L"
)

type Cake struct {
	Ingredients     []string
	ShakingMinuts   int8
	MoldSize        Size `validate:"oneof=S M L"`
	OvenTemperature float32
}

func SetCakeBuilder() *Cake {
	return &Cake{}
}

func (c *Cake) validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *Cake) MixIngredients() {

	sort.Strings(c.Ingredients)
	fmt.Printf("estoyyy %v", c)
}

func (c *Cake) ShakeMixture() {
	time.Sleep(time.Duration(c.ShakingMinuts) * time.Second)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(c.Ingredients),
		func(i, j int) {
			c.Ingredients[i], c.Ingredients[j] = c.Ingredients[j], c.Ingredients[i]
		},
	)
}

func (c *Cake) AddMixtureToMold() error {
	err := c.validate()

	if err != nil {
		ms := fmt.Sprintf(
			"Incorrect mold size choose between Small: %s Medium: %s Large: %s \n",
			SMALL, MEDIUM, LARGE,
		)
		panic(errors.New(ms))
	}
	return nil
}

func (c *Cake) CookInOven() {
	temperatureFtoC := (c.OvenTemperature - 32) * 5 / 9
	c.OvenTemperature = temperatureFtoC
}

func (c *Cake) GetCake() *Cake {
	return c
}

func LargeChocolateCakeBuilder() CakeBuilder {
	builder := SetCakeBuilder()
	builder.Ingredients = []string{"sugar", "butter", "eggs", "chocolate"}
	builder.MoldSize = LARGE
	builder.OvenTemperature = 70
	builder.ShakingMinuts = 3
	return builder
}

func SmallChocolateCakeBuilder() CakeBuilder {
	builder := SetCakeBuilder()
	builder.Ingredients = []string{"sugar", "butter", "eggs", "chocolate"}
	builder.MoldSize = SMALL
	builder.OvenTemperature = 50
	builder.ShakingMinuts = 2
	return builder
}
