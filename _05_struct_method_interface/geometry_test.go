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
	testsArea := []struct {
		form Form
		want float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range testsArea {
		got := tt.form.Area()
		if tt.want != got {
			t.Errorf(" queria '%.2f', obteve '%.2f'", tt.want, got)
		}
	}

}
