package goutil

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var upperCase = cases.Upper(language.English)

func HasUpperCase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}

	return false
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func SnakeToCamel(snake string, capitalize bool) string {
	parts := strings.Split(snake, "_")
	camel := strings.ToLower(parts[0])

	if len(parts) == 1 {
		return Capitalize(parts[0])
	}

	for _, part := range parts[1:] {
		if len(part) == 0 {
			continue
		}

		if len(part) > 1 {
			lower := strings.ToLower(part)
			camel += upperCase.String(lower[:1]) + lower[1:]
		} else {
			camel += Capitalize(part)
		}
	}

	if capitalize {
		return Capitalize(camel)
	}

	return camel
}

func CamelToSnake(camel string) string {
	var builder strings.Builder
	builder.Grow(len(camel)) // Preallocate memory

	var start int // Keep track of the start of the current lowercase substring

	for i, r := range camel {
		if unicode.IsUpper(r) {
			// Write the accumulated lowercase slice to the builder
			if start < i {
				builder.WriteString(camel[start:i])
			}

			// Add an underscore before it unless it's the first letter
			if i > 0 {
				builder.WriteString("_")
			}

			// Write the current upper-case letter as lowercase
			builder.WriteString(strings.ToLower(string(r)))

			// Move the start to the next character
			start = i + 1
		}
	}

	// Append the remaining lowercase characters after the last upper-case letter
	if start < len(camel) {
		builder.WriteString(camel[start:])
	}

	return builder.String()
}
