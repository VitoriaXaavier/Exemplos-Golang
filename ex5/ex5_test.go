package ex5

import (
	"testing"
	"math"
	"github.com/stretchr/testify/assert"
)

func TestForca(t * testing.T) {
	altura := 4.0
	diametro := 0.5 
	gama := 50.0

	forca := ForcaExercida(float64(altura),float64(diametro), float64(gama))

	pi := 3.14159
	exp2 := math.Pow(diametro, 2)

	f := (pi * gama * exp2 * altura) / 4
	assert.Equal(t, f, forca )

}