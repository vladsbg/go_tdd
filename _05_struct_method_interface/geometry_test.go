package _05_struct_method_interface

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("calculo perimetro", func(t *testing.T) {
		want := 40.0
		got := Perimeter(10.0, 10.0)

		if want != got {
			t.Errorf(" queria '%.2f', obteve '%.2f'", want, got)
		}
	})

}

func TestArea(t *testing.T) {
	t.Run("calculo area", func(t *testing.T) {
		want := 72.0
		got := Area(12.0, 6.0)

		if want != got {
			t.Errorf(" queria '%.2f', obteve '%.2f'", want, got)
		}
	})

}
