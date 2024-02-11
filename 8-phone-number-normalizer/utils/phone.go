package utils

import (
	"strings"
)

func RemoveUnwantedRunes(phoneNumbers []string) []string {
	var result []string

	for _, phone := range phoneNumbers {
		// Define a mapping function to remove specified characters
		// TODO: Make sanitization more robust; hardcoded checks won't scale
		removeChars := func(r rune) rune {
			switch r {
			case '-', '(', ')', ' ':
				return -1 // Remove the character
			default:
				return r // Keep other characters unchanged
			}
		}

		// Apply the mapping function to remove specified characters
		sanitizedPhone := strings.Map(removeChars, phone)
		result = append(result, sanitizedPhone)
	}

	return result
}
