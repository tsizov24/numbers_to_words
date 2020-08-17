package numbers_to_words

import (
	"strconv"
	"strings"
)

var (
	units = []string{"один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	units2 = []string{"одна", "две"}
	dozens = []string{"десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	dozens2 = []string{"одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	hundreds = []string{"сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
	nums = []string{"тысяч", "миллион", "миллиард", "триллион", "квадриллион", "квинтиллион", "секстиллион", "септиллион", "октиллион", "нониллион", "дециллион"}
)

func Ru(num string) (string, error) {
	if num == "0" {
		return "ноль", nil
	}
	if err := isValidNumber(num); err != nil {
		return "", err
	}
	if len(num) > len(nums) * 3 + 3 {
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
			sb.writeStringWithSpaces(hundreds[n1 - 1])
		}
		if n2 == 1 && n3 > 0 {
			sb.writeStringWithSpaces(dozens2[n3 - 1])
		} else {
			if n2 > 0 {
				sb.writeStringWithSpaces(dozens[n2 - 1])
			}
			if n3 > 0 {
				if cnt == 1 && (n3 == 1 || n3 == 2) {
					sb.writeStringWithSpaces(units2[n3 - 1])
				} else {
					sb.writeStringWithSpaces(units[n3 - 1])
				}
			}
		}
		if cnt > 0 && n > 0 {
			sb.WriteString(nums[cnt - 1])
			if cnt == 1 {
				if n2 != 1 {
					if n3 == 1 {
						sb.WriteString("а")
					} else if n3 >= 2 && n3 <= 4 {
						sb.WriteString("и")
					}
				}
			} else {
				if n2 == 1 || n3 > 4 {
					sb.WriteString("ов")
				} else if n3 >= 2 || n3 <= 4 {
					sb.WriteString("а")
				}
			}
		}
		sb.WriteString(" ")
		res.addAllAndReset(&sb)
		cnt++
	}
	return strings.Trim(res.String(), " "), nil
}
