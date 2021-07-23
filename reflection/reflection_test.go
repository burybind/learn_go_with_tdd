package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Occupation struct {
		Company, Title string
	}
	testCases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Brendan"},
			[]string{"Brendan"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Brendan", "Highland"},
			[]string{"Brendan", "Highland"},
		},
		{
			"Struct with one string field and one int field",
			struct {
				Name string
				Age  int
			}{"Brendan", 33},
			[]string{"Brendan"},
		},
		{
			"Struct with nested struct field",
			struct {
				Name string
				Occupation
				Age int
			}{"Brendan", Occupation{"Weave", "Software Engineer"}, 33},
			[]string{"Brendan", "Weave", "Software Engineer"},
		},
		{
			"Pointers to things",
			&struct {
				Name string
				Occupation
				Age int
			}{"Brendan", Occupation{"Weave", "Software Engineer"}, 33},
			[]string{"Brendan", "Weave", "Software Engineer"},
		},
		{
			"Slices",
			[]Occupation{
				{"Weave", "Software Engineer"},
				{"Weave", "Soda Drinker"},
			},
			[]string{"Weave", "Software Engineer", "Weave", "Soda Drinker"},
		},
		{
			"Arrays",
			[2]Occupation{
				{"Weave", "Software Engineer"},
				{"Weave", "Soda Drinker"},
			},
			[]string{"Weave", "Software Engineer", "Weave", "Soda Drinker"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Ice Cream": "Vanilla",
			"Color":     "Green",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Vanilla")
		assertContains(t, got, "Green")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Occupation)

		go func() {
			aChannel <- Occupation{"Weave", "Software Engineer"}
			aChannel <- Occupation{"Weave", "Soda Drinker"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Weave", "Software Engineer", "Weave", "Soda Drinker"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Occupation, Occupation) {
			return Occupation{"Weave", "Software Engineer"}, Occupation{"Weave", "Soda Drinker"}
		}

		var got []string
		want := []string{"Weave", "Software Engineer", "Weave", "Soda Drinker"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
