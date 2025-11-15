package utils

import (
	"strings"
	"unicode"
)

func CalculateAccuracy(targetText, transcribedText string) float64 {
	target := normalizeText(targetText)
	transcribed := normalizeText(transcribedText)

	if len(target) == 0 && len(transcribed) == 0 {
		return 100.0
	}

	if len(target) == 0 || len(transcribed) == 0 {
		return 0.0
	}

	distance := levenshteinDistance(target, transcribed)
	maxLen := len(target)
	if len(transcribed) > maxLen {
		maxLen = len(transcribed)
	}

	accuracy := 100.0 * (1.0 - float64(distance)/float64(maxLen))
	if accuracy < 0 {
		accuracy = 0
	}
	if accuracy > 100 {
		accuracy = 100
	}

	return accuracy
}

// normalizeText - normalize text untuk comparison
func normalizeText(text string) string {
	// Lowercase
	text = strings.ToLower(text)

	// Remove punctuation
	text = removePunctuation(text)

	// Trim extra spaces
	text = strings.TrimSpace(text)

	// Remove multiple spaces
	text = strings.Join(strings.Fields(text), " ")

	return text
}

// removePunctuation - hapus tanda baca
func removePunctuation(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, text)
}

// levenshteinDistance - hitung edit distance antara dua string
func levenshteinDistance(a, b string) int {
	lenA := len(a)
	lenB := len(b)

	matrix := make([][]int, lenA+1)
	for i := range matrix {
		matrix[i] = make([]int, lenB+1)
	}

	for i := 0; i <= lenA; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lenB; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[lenA][lenB]
}

// min - helper function
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
