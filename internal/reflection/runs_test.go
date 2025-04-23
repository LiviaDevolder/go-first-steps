package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestRuns(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct as a string",
			struct {
				Name string
			}{"Livia"},
			[]string{"Livia"},
		},
		{
			"Struct as two string fields",
			struct {
				Name string
				City string
			}{"Livia", "Niterói"},
			[]string{"Livia", "Niterói"},
		},
		{
			"Struct without string field",
			struct {
				Name string
				Age  int
			}{"Livia", 22},
			[]string{"Livia"},
		},
		{
			"Nested fields",
			Person{"Livia", Profile{22, "Niterói"}},
			[]string{"Livia", "Niterói"},
		},
		{
			"Pointers to things",
			&Person{
				"Livia",
				Profile{22, "Niterói"},
			},
			[]string{"Livia", "Niterói"},
		},
		{
			"Slices",
			[]Profile{
				{22, "Niterói"},
				{23, "Porto Alegre"},
			},
			[]string{"Niterói", "Porto Alegre"},
		},
		{
			"Arrays",
			[2]Profile{
				{22, "Niterói"},
				{23, "Porto Alegre"},
			},
			[]string{"Niterói", "Porto Alegre"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			// Arrange
			var result []string

			// Act
			Runs(test.Input, func(input string) {
				result = append(result, input)
			})

			// Assert
			if !reflect.DeepEqual(result, test.ExpectedCalls) {
				t.Errorf("result %v, expect %v", result, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		mapA := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var result []string

		Runs(mapA, func(input string) {
			result = append(result, input)
		})

		verifyIfContains(t, result, "Bar")
		verifyIfContains(t, result, "Boz")
	})
}

func verifyIfContains(t *testing.T, haystack []string, needle string) {
	contain := false

	for _, x := range haystack {
		if x == needle {
			contain = true
		}
	}

	if !contain {
		t.Errorf("expect that %+v has %s, but didnt", haystack, needle)
	}
}
