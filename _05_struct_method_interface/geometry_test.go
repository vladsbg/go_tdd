package _05_struct_method_interface

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("calculo perimetro", func(t *testing.T) {
		want := 40.0
		got := Perimeter(Rectangle{10.0, 10.0})

		if want != got {
			t.Errorf(" queria '%.2f', obteve '%.2f'", want, got)
		}
	})

}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, form Form, want float64) {
		t.Helper()
		got := form.Area()

		if want != got {
			t.Errorf(" queria '%.2f', obteve '%.2f'", want, got)
		}
	}

	t.Run("calcula area retangulo", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("calcula area círculo", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
