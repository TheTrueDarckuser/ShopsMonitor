package pystrings

import "strings"

// String adds pipeable methods to string
type String string

// Split string by separator
func (s String) Split(sep string) (pieces []String) {
	temp := strings.Split(string(s), sep)
	for _, t := range temp {
		pieces = append(pieces, String(t))
	}

	return pieces
}

// Trim trims string
func (s String) Trim(cutset string) String {
	return String(strings.Trim(string(s), cutset))
}

// Lower lowers string
func (s String) Lower() String {
	return String(strings.ToLower(string(s)))
}
