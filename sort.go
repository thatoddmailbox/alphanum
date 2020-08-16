package alphanum

import (
	"unicode"
	"unicode/utf8"
)

func breakIntoChunks(str string) ([]string, []bool) {
	chunks := []string{}
	numeric := []bool{}

	currentChunk := ""
	currentChunkNumeric := false
	for {
		r, s := utf8.DecodeRuneInString(str)
		if r == utf8.RuneError {
			break
		}

		isDigit := unicode.IsDigit(r)
		if len(currentChunk) > 0 {
			if isDigit != currentChunkNumeric {
				// we've reached the end of the last chunk
				// save it
				chunks = append(chunks, currentChunk)
				numeric = append(numeric, currentChunkNumeric)

				// and start a new one
				currentChunk = string(r)
				currentChunkNumeric = isDigit
			} else {
				// continue the current chunk
				currentChunk += string(r)
			}
		} else {
			// starting a new chunk
			currentChunk = string(r)
			currentChunkNumeric = isDigit
		}

		str = str[s:]
	}

	if len(currentChunk) > 0 {
		// we still have a leftover chunk, add it
		chunks = append(chunks, currentChunk)
		numeric = append(numeric, currentChunkNumeric)
	}

	return chunks, numeric
}

func lessChunks(aChunks []string, aNumeric []bool, bChunks []string, bNumeric []bool) bool {
	aChunkCount := len(aChunks)
	bChunkCount := len(bChunks)

	smallerChunkCount := aChunkCount
	if bChunkCount < aChunkCount {
		smallerChunkCount = bChunkCount
	}

	for i := 0; i < smallerChunkCount; i++ {
		aChunk := aChunks[i]
		bChunk := bChunks[i]

		if aNumeric[i] && bNumeric[i] {
			// compare as numbers
			// as described on the website, we shouldn't convert these to ints, since we might run into trouble with larger numbers
			if len(aChunk) != len(bChunk) {
				// larger length == bigger number
				return (len(aChunk) < len(bChunk))
			}

			for j := 0; j < len(aChunk); j++ {
				aDigit := aChunk[j] - '0'
				bDigit := bChunk[j] - '0'
				if aDigit != bDigit {
					return (aDigit < bDigit)
				}
			}
		} else {
			// compare as strings
			if aChunk == bChunk {
				continue
			}

			return (aChunk < bChunk)
		}
	}

	if aChunkCount < bChunkCount {
		return true
	}

	return false
}

// Less reports whether string A should be before string B.
func Less(a string, b string) bool {
	aChunks, aNumeric := breakIntoChunks(a)
	bChunks, bNumeric := breakIntoChunks(b)
	return lessChunks(aChunks, aNumeric, bChunks, bNumeric)
}
