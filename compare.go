
package main

import (
	"fmt"
	"strings"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

func main() {
	// String with precomposed character (NFC)
	s1 := "résumé" 

	// String with decomposed character (NFD)
	s2 := "résumé" 

	fmt.Printf("s1: %s (NFC form)\n", s1)
	fmt.Printf("s2: %s (NFD form)\n", s2)

	// Direct comparison will return false
	fmt.Printf("Direct comparison (s1 == s2): %t\n", s1 == s2)

	// Normalize both strings to NFC for comparison
	normalizedS1 := norm.NFC.String(s1)
	normalizedS2 := norm.NFC.String(s2)

	fmt.Printf("Normalized s1 (NFC): %s\n", normalizedS1)
	fmt.Printf("Normalized s2 (NFC): %s\n", normalizedS2)

	// Compare normalized strings
	fmt.Printf("Normalized comparison (normalizedS1 == normalizedS2): %t\n", normalizedS1 == normalizedS2)

	// Using strings.Compare for lexicographical comparison after normalization
	comparisonResult := strings.Compare(normalizedS1, normalizedS2)
	fmt.Printf("strings.Compare on normalized strings: %d (0 means equal)\n", comparisonResult)


    str1 := "Café"
	str2 := "Cafe"
	str3 := "café"

	// Create a collator that ignores case and diacritics
	c := collate.New(language.Und, collate.IgnoreCase, collate.IgnoreDiacritics)

	// Compare str1 and str2
	result1 := c.CompareString(str1, str2)
	fmt.Printf("'%s' vs '%s': %d\n", str1, str2, result1) // Expected: 0 (equal)

	// Compare str1 and str3
	result2 := c.CompareString(str1, str3)
	fmt.Printf("'%s' vs '%s': %d\n", str1, str3, result2) // Expected: 0 (equal, due to IgnoreCase)
}