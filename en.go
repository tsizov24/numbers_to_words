package numbers_to_words

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	enUnits   = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	enDozens  = []string{"ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	enDozens2 = []string{"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	enNums    = []string{"thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion"}
)

func En(number interface{}) (string, error) {
	num := fmt.Sprintf("%v", number)
	if num == "0" {
		return "zero", nil
	}
	if err := isValidNumber(num); err != nil {
		return "", err
	}
	if len(num) > len(enNums) * 3 + 3 {
		return "", tooBigNumber
	}
	res := stringBuilder{}
	sb := stringBuilder{}
	pos := len(num)
	cnt := 0
	for pos > 0 {
		n, _ := strconv.Atoi(num[max(0, pos - 3):pos])
		n1 := n / 100
		n2 := n / 10 % 10
		n3 := n % 100 % 10
		pos -= 3
		if n1 > 0 {
			sb.writeStringWithSpaces(enUnits[n1 - 1])
			sb.writeStringWithSpaces("hundred")
		}
		if n2 == 1 && n3 > 0 {
			sb.writeStringWithSpaces(enDozens2[n3 - 1])
		} else {
			if n2 > 0 {
				if n3 > 0 {
					_, _ = sb.WriteString(enDozens[n2 - 1])
					_, _ = sb.WriteString("-")
				} else {
					sb.writeStringWithSpaces(enDozens[n2 - 1])
				}
			}
			if n3 > 0 {
				sb.writeStringWithSpaces(enUnits[n3 - 1])
			}
		}
		if cnt > 0 && n > 0 {
			sb.writeStringWithSpaces(enNums[cnt - 1])
		}
		res.addAllAndReset(&sb)
		cnt++
	}
	return strings.Trim(res.String(), " "), nil
}
