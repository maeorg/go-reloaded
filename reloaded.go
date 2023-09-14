package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		println("File name missing")
		return
	} else if len(os.Args) > 3 {
		println("Too many arguments")
		return
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		println("Error: ", err)
	}

	input := string(file)

	temp := FormatText(input)
	result := FormatText(temp)

	f, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(result)
	if err != nil {
		log.Fatal(err)
	}
}

func FormatText(input string) string {
	words := strings.Split(input, " ")
	quatationMarkCounter := 0
	for i := 1; i < len(words); i++ {
		word := words[i]

		if strings.ToLower(words[i]) == "a" {
			c := strings.ToLower(string(words[i+1][0]))
			if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" || c == "h" {
				if words[i] == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		}

		switch word {
		case "(hex)":
			words[i-1] = ConvertBase(strings.ToUpper((words[i-1])), "0123456789ABCDEF", "0123456789")
			words = append(words[:i], words[i+1:]...)
		case "(bin)":
			words[i-1] = ConvertBase(string(words[i-1]), "01", "0123456789")
			words = append(words[:i], words[i+1:]...)
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(cap)":
			temp := strings.ToUpper(string(words[i-1][0]))
			temp += strings.ToLower(words[i-1][1:])
			words[i-1] = temp
			words = append(words[:i], words[i+1:]...)
		case "(up,":
			temp := words[i+1]
			number := Atoi(temp[:len(temp)-1])
			count := 0
			for n := 1; n <= number; n++ {
				words[i-n] = strings.ToUpper(words[i-n])
				if count < 2 {
					words = append(words[:i], words[i+1:]...)
					count++
				}
			}
		case "(low,":
			temp := words[i+1]
			number := Atoi(temp[:len(temp)-1])
			count := 0
			for n := 1; n <= number; n++ {
				words[i-n] = strings.ToLower(words[i-n])
				if count < 2 {
					words = append(words[:i], words[i+1:]...)
					count++
				}
			}
		case "(cap,":
			next := words[i+1]
			number := Atoi(next[:len(next)-1])
			count := 0
			for n := 1; n <= number; n++ {
				temp := strings.ToUpper(string(words[i-n][0]))
				temp += strings.ToLower(words[i-n][1:])
				words[i-n] = temp
				if count < 2 {
					words = append(words[:i], words[i+1:]...)
					count++
				}
			}
		case "'":
			quatationMarkCounter++
			if quatationMarkCounter == 1 {
				words[i+1] = "'" + words[i+1]
			} else if quatationMarkCounter == 2 {
				words[i-1] = words[i-1] + "'"
				quatationMarkCounter = 0
			}
			words = append(words[:i], words[i+1:]...)
		}

		if len(words) > i+1 {
			word = words[i]
		}
		matched, _ := regexp.MatchString(`\A[.,!?:;]+`, word)
		re := regexp.MustCompile(`\A[.,!?:;]+`)
		punctuation := re.FindString(word)
		if matched {
			// if the punctuation and some letters are in the same word split, also add the rest of the word to the sentence
			if len(word) > len(punctuation) {
				words[i-1] = words[i-1] + punctuation + " " + words[i][len(punctuation):]
			} else { // if the word contains only punctuations, only add the punctuations to the end of previous word
				words[i-1] = words[i-1] + punctuation
			}
			words = append(words[:i], words[i+1:]...)
		}
				
	}

	result := string(strings.Join(words, " "))
	return result
}


// ConvertBase("328838", "0123456789", "0123456789ABCDEF") == "" instead of "50486"
func ConvertBase(nbr, baseFrom, baseTo string) string {
	return PrintNbrBase2(AtoiBase(nbr, baseFrom), baseTo)
}

func PrintNbrBase2(nbr int, base string) string {
	digits := []rune{}
	for nbr != 0 {
		mod := nbr % len(base)
		digits = append(digits, rune(base[mod]))
		nbr /= len(base)
	}
	if len(digits) == 0 {
		digits = append(digits, '0')
	}
	newDigits := []rune{}
	for i := len(digits) - 1; i >= 0; i-- {
		newDigits = append(newDigits, digits[i])
	}
	return string(newDigits)
}

func AtoiBase(s string, base string) int {
	mp := make(map[rune]int)
	for i, c := range base {
		_, ok := mp[c]
		if c == '+' || c == '-' || len(base) < 2 || ok {
			return 0
		}
		mp[c] = i
	}

	result := 0
	for i, n := range s {
		index := mp[n]
		power := RecursivePower(len(base), len(s)-1-i)
		result += index * power
	}

	return result
}

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	if power > 0 {
		return nb * RecursivePower(nb, power-1)
	}
	return 1
}

func Atoi(s string) int {
	runeString := []rune(s)
	number := 0
	for i, j := len(runeString)-1, 1; i >= 0; i, j = i-1, j*10 {
		if i == 0 && runeString[0] == '+' {
			return number
		} else if i == 0 && runeString[0] == '-' {
			return number * -1
		}
		if runeString[i] < '0' || runeString[i] > '9' {
			return 0
		}
		number += int(runeString[i]-'0') * j
	}
	return number
}
