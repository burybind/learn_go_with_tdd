package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		desc  string
		shape Shape
		want  float64
	}{
		{
			desc:  "area on rectangle",
			shape: Rectangle{10.0, 10.0},
			want:  100.0,
		},
		{
			desc:  "area on circle",
			shape: Circle{10.0},
			want:  314.1592653589793,
		},
		{
			desc:  "area on triangle",
			shape: Triangle{12, 6},
			want:  36,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.shape.Area()
			if got != tC.want {
				t.Errorf("%#v got %g want %g", tC.shape, got, tC.want)
			}
		})
	}
}
