package struct_method_interface

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
		name string
		form Form
		want float64
	}{
		{name: "retangulo", form: Rectangle{Height: 6, Width: 12}, want: 72.0},
		{name: "circulo", form: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangulo", form: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.form.Area()
			if tt.want != got {
				t.Errorf("%#v queria '%.2f', obteve '%.2f'", tt.name, tt.want, got)
			}
		})
	}
}
