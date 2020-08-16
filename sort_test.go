package alphanum

import "testing"

func testSingleChunk(t *testing.T, s string, chunks []string, numerics []bool) {
	resultChunks, resultNumerics := breakIntoChunks(s)
	pass := true
	if len(resultChunks) == len(chunks) && len(resultNumerics) == len(numerics) {
		for i, resultChunk := range resultChunks {
			if chunks[i] != resultChunk {
				pass = false
				break
			}
		}

		if pass {
			for i, resultNumeric := range resultNumerics {
				if numerics[i] != resultNumeric {
					pass = false
					break
				}
			}
		}
	} else {
		pass = false
	}

	if !pass {
		t.Errorf(
			"breaking '%s' into chunks:\n\tchunks: got %v, expected %v\n\tnumerics: got %v, expected %v",
			s,
			resultChunks, chunks,
			resultNumerics, numerics,
		)
	}
}

func TestChunks(t *testing.T) {
	testSingleChunk(t, "z123", []string{"z", "123"}, []bool{false, true})
	testSingleChunk(t, "z1y2", []string{"z", "1", "y", "2"}, []bool{false, true, false, true})
	testSingleChunk(t, "5*10e-2", []string{"5", "*", "10", "e-", "2"}, []bool{true, false, true, false, true})
	testSingleChunk(t, "1000X Radonius Maximus", []string{"1000", "X Radonius Maximus"}, []bool{true, false})
	testSingleChunk(t, "Callisto Morphamax 6000 SE", []string{"Callisto Morphamax ", "6000", " SE"}, []bool{false, true, false})
	testSingleChunk(t, "Callisto Morphamax 6000 SE2", []string{"Callisto Morphamax ", "6000", " SE", "2"}, []bool{false, true, false, true})
}

func testSingleLess(t *testing.T, a string, b string) {
	result := Less(a, b)
	if !result {
		t.Errorf("Less('%s', '%s'): got false, expected true", a, b)
	}

	flippedResult := Less(b, a)
	if flippedResult {
		t.Errorf("Less('%s', '%s'): got true, expected false", b, a)
	}
}

func TestLess(t *testing.T) {
	testSingleLess(t, "z1", "z2")
	testSingleLess(t, "z1", "z100")
	testSingleLess(t, "z2", "z10")
	testSingleLess(t, "Callisto Morphamax 6000 SE", "Callisto Morphamax 6000 SE2")
	testSingleLess(t, "Allegia 50 Clasteron", "Allegia 500 Clasteron")
	testSingleLess(t, "Callisto Morphamax 7000", "Xiph Xlater 5")
}
