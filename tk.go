package numbers_to_words

import (
	"strconv"
	"strings"
)

var (
	tkUnits  = []string{"bir", "iki", "üç", "dört", "bäş", "alty", "ýedi", "sekiz", "dokuz"}
	tkDozens = []string{"on", "ýigrimi", "otuz", "kyrk", "elli", "altmyş", "yetmiş", "segsen", "togsan"}
	tkNums   = []string{"müň", "million", "milliard", "trillion", "kwadrillion", "kwintillion", "sekstillion", "septillion", "oktillion", "nonillion", "desillion"}
)

func Tk(num string) (string, error) {
	if num == "0" {
		return "nol", nil
	}
	if err := isValidNumber(num); err != nil {
		return "", err
	}
	if len(num) > len(tkNums) * 3 + 3 {
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
			sb.writeStringWithSpaces(tkUnits[n1 - 1])
			sb.writeStringWithSpaces("yüz")
		}
		if n2 > 0 {
			sb.writeStringWithSpaces(tkDozens[n2 - 1])
		}
		if n3 > 0 {
			sb.writeStringWithSpaces(tkUnits[n3 - 1])
		}
		if cnt > 0 && n > 0 {
			sb.writeStringWithSpaces(tkNums[cnt - 1])
		}
		res.addAllAndReset(&sb)
		cnt++
	}
	return strings.Trim(res.String(), " "), nil
}
