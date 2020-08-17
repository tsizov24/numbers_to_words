package numbers_to_words

import (
	"errors"
	"math"
	"regexp"
	"strings"
)

type stringBuilder struct {
	strings.Builder
}

var (
	tooBigNumber = errors.New("too big number")
	wrongNumber = errors.New("wrong number format")
)

// Checks validity of number. Returns nil if number is prime.
func isValidNumber(num string) error { //!!!!!!!!!!!!!
	valid := regexp.MustCompile(`^[1-9][0-9]*$`)
	if valid.MatchString(num) {
		return nil
	}
	return wrongNumber
}

// Returns maximum number
func max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}

func (sb *stringBuilder) writeStringWithSpaces(s string) {
	sb.WriteString(s)
	sb.WriteString(" ")
}

func (sb *stringBuilder) addAllAndReset(sb2 * stringBuilder) {
	sb2.WriteString(sb.String())
	sb.Reset()
	sb.WriteString(sb2.String())
	sb2.Reset()
}