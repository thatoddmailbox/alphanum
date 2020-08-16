package alphanum

import "testing"

func TestReadme(t *testing.T) {
	// Usage -> sort.Interface
	testSingleSliceSort(t, []string{
		"20X Radonius Prime",
		"Allegia 50 Clasteron",
		"20X Radonius",
		"Xiph Xlater 3",
		"Xiph Xlater 1",
		"Xiph Xlater 58",
		"Xiph Xlater 12",
		"Xiph Xlater 10",
		"Xiph Xlater 2",
		"Callisto Morphamax 6000 SE",
		"Xiph Xlater 11",
	}, []string{
		"20X Radonius",
		"20X Radonius Prime",
		"Allegia 50 Clasteron",
		"Callisto Morphamax 6000 SE",
		"Xiph Xlater 1",
		"Xiph Xlater 2",
		"Xiph Xlater 3",
		"Xiph Xlater 10",
		"Xiph Xlater 11",
		"Xiph Xlater 12",
		"Xiph Xlater 58",
	})

	// Usage -> Less
	testSingleLess(t, "File 2", "File 10")
	testSingleLess(t, "File 10", "File 17")
	testSingleLess(t, "Allegia 50 Clasteron", "Allegia 500 Clasteron")
	testSingleLess(t, "Callisto Morphamax 6000 SE", "Callisto Morphamax 6000 SE2")
}
