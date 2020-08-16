package alphanum

import (
	"sort"
	"testing"
)

func testSingleSliceSort(t *testing.T, input []string, expected []string) {
	result := StringSlice(append([]string{}, input...))
	sort.Sort(result)

	if len(result) != len(expected) {
		t.Errorf("sort %v: expected %v, got %v", input, expected, result)
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("sort %v: expected %v, got %v", input, expected, result)
		}
	}
}

func TestSliceSort(t *testing.T) {
	// smaller examples

	testSingleSliceSort(t, []string{
		"z17.doc", "z1.doc", "z2.doc", "z35.doc",
	}, []string{
		"z1.doc", "z2.doc", "z17.doc", "z35.doc",
	})
	testSingleSliceSort(t, []string{
		"Callisto Morphamax 7000",
		"Allegia 50 Clasteron",
		"Xiph Xlater 5",
		"1000X Radonius Maximus",
	}, []string{
		"1000X Radonius Maximus",
		"Allegia 50 Clasteron",
		"Callisto Morphamax 7000",
		"Xiph Xlater 5",
	})

	// full examples from webpage

	testSingleSliceSort(t, []string{
		"z5.doc",
		"z16.doc",
		"z6.doc",
		"z11.doc",
		"z10.doc",
		"z4.doc",
		"z100.doc",
		"z9.doc",
		"z20.doc",
		"z3.doc",
		"z15.doc",
		"z13.doc",
		"z14.doc",
		"z102.doc",
		"z1.doc",
		"z12.doc",
		"z101.doc",
		"z18.doc",
		"z17.doc",
		"z7.doc",
		"z19.doc",
		"z8.doc",
		"z2.doc",
	}, []string{
		"z1.doc",
		"z2.doc",
		"z3.doc",
		"z4.doc",
		"z5.doc",
		"z6.doc",
		"z7.doc",
		"z8.doc",
		"z9.doc",
		"z10.doc",
		"z11.doc",
		"z12.doc",
		"z13.doc",
		"z14.doc",
		"z15.doc",
		"z16.doc",
		"z17.doc",
		"z18.doc",
		"z19.doc",
		"z20.doc",
		"z100.doc",
		"z101.doc",
		"z102.doc",
	})
	testSingleSliceSort(t, []string{
		"Alpha 100",
		"Alpha 200",
		"10X Radonius",
		"Callisto Morphamax 500",
		"Xiph Xlater 500",
		"Xiph Xlater 50",
		"Xiph Xlater 5000",
		"Xiph Xlater 300",
		"Alpha 2A-8000",
		"Allegia 6R Clasteron",
		"30X Radonius",
		"40X Radonius",
		"Xiph Xlater 10000",
		"Callisto Morphamax 6000 SE2",
		"Alpha 2A",
		"Allegia 50 Clasteron",
		"Xiph Xlater 58",
		"20X Radonius Prime",
		"Callisto Morphamax 6000 SE",
		"20X Radonius",
		"Xiph Xlater 2000",
		"200X Radonius",
		"Callisto Morphamax",
		"Allegia 500 Clasteron",
		"Callisto Morphamax 700",
		"Alpha 2",
		"Xiph Xlater 40",
		"Xiph Xlater 5",
		"Allegia 50B Clasteron",
		"Alpha 2A-900",
		"Allegia 51 Clasteron",
		"Callisto Morphamax 5000",
		"1000X Radonius Maximus",
		"Callisto Morphamax 7000",
		"Callisto Morphamax 600",
	}, []string{
		"10X Radonius",
		"20X Radonius",
		"20X Radonius Prime",
		"30X Radonius",
		"40X Radonius",
		"200X Radonius",
		"1000X Radonius Maximus",
		"Allegia 6R Clasteron",
		"Allegia 50 Clasteron",
		"Allegia 50B Clasteron",
		"Allegia 51 Clasteron",
		"Allegia 500 Clasteron",
		"Alpha 2",
		"Alpha 2A",
		"Alpha 2A-900",
		"Alpha 2A-8000",
		"Alpha 100",
		"Alpha 200",
		"Callisto Morphamax",
		"Callisto Morphamax 500",
		"Callisto Morphamax 600",
		"Callisto Morphamax 700",
		"Callisto Morphamax 5000",
		"Callisto Morphamax 6000 SE",
		"Callisto Morphamax 6000 SE2",
		"Callisto Morphamax 7000",
		"Xiph Xlater 5",
		"Xiph Xlater 40",
		"Xiph Xlater 50",
		"Xiph Xlater 58",
		"Xiph Xlater 300",
		"Xiph Xlater 500",
		"Xiph Xlater 2000",
		"Xiph Xlater 5000",
		"Xiph Xlater 10000",
	})
}
