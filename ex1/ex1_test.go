package ex1

import (
	"testing"
)

func TestArea(t *testing.T) {

	area := Area(20, 10)
	if area != 100 {
	t.Error("Erro no teste")
	}

	area2 := Area(2,4)
	if area2 != 4 {
		t.Error("Erro no teste 2")
	}
}

