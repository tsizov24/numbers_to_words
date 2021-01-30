package numbers_to_words

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ruUnits    = []string{"один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	ruUnits2   = []string{"одна", "две"}
	ruDozens   = []string{"десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	ruDozens2  = []string{"одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	ruHundreds = []string{"сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
	ruNums     = []string{"тысяч", "миллион", "миллиард", "триллион", "квадриллион", "квинтиллион", "секстиллион", "септиллион", "октиллион", "нониллион", "дециллион"}
)

func Ru(number interface{}) (string, error) {
	num := fmt.Sprintf("%v", number)
	if num == "0" {
		return "ноль", nil
	}
	if err := isValidNumber(num); err != nil {
		return "", err
	}
	if len(num) > len(ruNums) * 3 + 3 {
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
			sb.writeStringWithSpaces(ruHundreds[n1 - 1])
		}
		if n2 == 1 && n3 > 0 {
			sb.writeStringWithSpaces(ruDozens2[n3 - 1])
		} else {
			if n2 > 0 {
				sb.writeStringWithSpaces(ruDozens[n2 - 1])
			}
			if n3 > 0 {
				if cnt == 1 && (n3 == 1 || n3 == 2) {
					sb.writeStringWithSpaces(ruUnits2[n3 - 1])
				} else {
					sb.writeStringWithSpaces(ruUnits[n3 - 1])
				}
			}
		}
		if cnt > 0 && n > 0 {
			_, _ = sb.WriteString(ruNums[cnt - 1])
			if cnt == 1 {
				if n2 != 1 {
					if n3 == 1 {
						_, _ = sb.WriteString("а")
					} else if n3 >= 2 && n3 <= 4 {
						_, _ = sb.WriteString("и")
					}
				}
			} else {
				if n2 == 1 || n3 > 4 {
					_, _ = sb.WriteString("ов")
				} else if n3 >= 2 || n3 <= 4 {
					_, _ = sb.WriteString("а")
				}
			}
		}
		_, _ = sb.WriteString(" ")
		res.addAllAndReset(&sb)
		cnt++
	}
	return strings.Trim(res.String(), " "), nil
}
