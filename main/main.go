package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 && TxtChecker(args[0]) && TxtChecker(args[1]) {
		arg1 := args[0]
		outfile := args[1]
		data, err1 := os.ReadFile(arg1)
		if err1 != nil {
			log.Fatal(err1)
		}
		str := string(data)
		whitespaces1 := regexp.MustCompile(`^[ \t]+`)
		whitespaces2 := regexp.MustCompile(`[ \t]+\z`)
		whitespaces3 := regexp.MustCompile(`[ \t]+\n`)
		whitespaces4 := regexp.MustCompile(`\n[ \t]+`)
		whitespaces := regexp.MustCompile(`[ \t][ \t]+`)

		str = string(whitespaces1.ReplaceAllString(str, ""))
		str = string(whitespaces2.ReplaceAllString(str, ""))
		str = string(whitespaces3.ReplaceAllString(str, "\n"))
		str = string(whitespaces.ReplaceAllString(str, " "))
		str = string(whitespaces4.ReplaceAllString(str, "\n"))

		arr := strings.Split(Comma_to_front(str), " ")
		new_arr := WordChecker(arr)
		new_arr = Vowel(new_arr)
		temp_str := strings.Join(new_arr, " ")
		new_str := whitespaces.ReplaceAllString(temp_str, " ")
		new_str = CommaToBack(new_str)
		new_str = FirstSpace(new_str)
		new_str = LastSpace(new_str)

		new_byte := []byte(new_str)
		err2 := os.WriteFile(outfile, new_byte, 0644)
		if err2 != nil {
			log.Fatal(err2)
		}
	} else {
		fmt.Println("Not correct input")
	}
}

var flag = true

func LastSpace(str string) string {
	if flag {
		if str[len(str)-1] == ' ' {
			str = str[0 : len(str)-1]
		}
	}
	return str
}

func FirstSpace(str string) string {
	if flag {
		if str[0] == ' ' {
			str = str[1:]
		}
	}
	return str
}

func TxtChecker(str string) bool {
	txtFlag := false
	if strings.HasSuffix(str, ".txt") {
		txtFlag = true
	}
	return txtFlag
}

func Comma_to_front(str string) string {
	for i := 1; i < len(str); i++ {
		if str[i] == '.' || str[i] == ',' || str[i] == '!' || str[i] == '?' || str[i] == ':' || str[i] == ';' || str[i] == ')' || str[i] == '(' {
			if str[i-1] != ' ' {
				str = strings.ReplaceAll(str, string(str[i]), string(" "+string((str[i]))))
			}
		}
	}
	for i := 0; i < len(str)-1; i++ {
		if str[i] == '.' || str[i] == ',' || str[i] == '!' || str[i] == '?' || str[i] == ':' || str[i] == ';' || str[i] == ')' {
			if str[i+1] != ' ' {
				str = strings.ReplaceAll(str, string(str[i]), string(string((str[i]))+" "))
			}
		}
	}
	str = strings.ReplaceAll(str, " '", " ' ")
	str = strings.ReplaceAll(str, "' ", " ' ")
	str = strings.ReplaceAll(str, "''", " '' ")
	str = strings.ReplaceAll(str, "\n", " \n ")
	str = strings.ReplaceAll(str, "  '  ", " ' ")
	str = strings.ReplaceAll(str, "  (", " (")
	str = strings.ReplaceAll(str, "  ", " ")

	return str
}

func CommaToBack(str string) string {
	for i := 0; i < len(str); i++ {
		if IsPunct(rune(str[i])) && str[i] != '(' && str[i] != '\'' {
			str = strings.ReplaceAll(str, " "+string(str[i]), string((str[i])))
		}
	}
	str = strings.ReplaceAll(str, " ' ' ", " '' ")
	str = strings.ReplaceAll(str, " \n ", "\n")
	str = strings.ReplaceAll(str, " \n", "\n")
	str = strings.ReplaceAll(str, "\n ", "\n")

	return str
}

func Vowel(arr []string) []string {
	var letter string
	vowel := false
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == "A" || arr[i] == "a" {
			letter = strings.ToLower(string(arr[i+1][0]))
			if letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u" || letter == "h" {
				vowel = true
			}
		}
		if arr[i] == "a" && vowel {
			arr[i] = "an"
		} else if arr[i] == "A" && vowel {
			arr[i] = "An"
		} else if arr[i] == "an" {
			arr[i] = "a"
		} else if arr[i] == "An" {
			arr[i] = "A"
		}
	}
	return arr
}

func WordChecker(arr []string) []string {
	if arr[0] == "(up" || arr[0] == "(low" || arr[0] == "(cap" || arr[0] == "(hex" || arr[0] == "(bin" {
		flag = false
	}
	var new_arr []string
	for i := 0; i < len(arr); i++ {
		if flag == true {
			if arr[i] == "(up" {
				new_arr = Up(arr, i)
			} else if arr[i] == "(low" {
				new_arr = Low(arr, i)
			} else if arr[i] == "(cap" {
				new_arr = Cap(arr, i)
			} else if arr[i] == "(hex" {
				new_arr = Hex(arr, i)
			} else if arr[i] == "(bin" {
				new_arr = Bin(arr, i)
			} else {
				new_arr = arr
			}
		} else {
			fmt.Println("ERROR")
			return nil
		}
	}
	new_arr = Quotation(new_arr)

	return new_arr
}

func Up(arr []string, i int) []string {
	if i-1 >= 0 && i+1 < len(arr) && string(arr[i]+arr[i+1]) == "(up)" {
		for j := 0; j < i; j++ {
			if i-1-j >= 0 && IsLetter(arr[i-1-j]) {
				arr[i-1-j] = strings.ToUpper(arr[i-1-j])
				break
			}
		}
		arr[i], arr[i+1] = " ", " "
	} else if i+3 < len(arr) && string(arr[i]+arr[i+1]) == "(up," {
		if IsNumber(arr[i+2]) {
			num := TypeChanger(arr[i+2])
			for j := 1; j <= num; j++ {
				if i >= num && i-j >= 0 {
					if IsLetter(arr[i-j]) && arr[i-j] != "\n" {
						arr[i-j] = strings.ToUpper(arr[i-j])
					} else {
						num++
					}
				} else {
					flag = false
					return nil
				}
			}
			arr[i], arr[i+1], arr[i+2], arr[i+3] = " ", " ", " ", " "
		}
	}
	return arr
}

func Low(arr []string, i int) []string {
	if i-1 >= 0 && i+1 < len(arr) && string(arr[i]+arr[i+1]) == "(low)" {
		for j := 0; j < i; j++ {
			if i-1-j >= 0 && IsLetter(arr[i-1-j]) {
				arr[i-1-j] = strings.ToLower(arr[i-1-j])
				break
			}
		}
		arr[i], arr[i+1] = " ", " "
	} else if i+3 < len(arr) && string(arr[i]+arr[i+1]) == "(low," {
		if IsNumber(arr[i+2]) {
			num := TypeChanger(arr[i+2])
			for j := 1; j <= num; j++ {
				if i >= num && i-j >= 0 {
					if IsLetter(arr[i-j]) {
						arr[i-j] = strings.ToLower(arr[i-j])
					} else {
						num++
					}
				} else {
					flag = false
					return nil
				}
			}
			arr[i], arr[i+1], arr[i+2], arr[i+3] = " ", " ", " ", " "
		}
	}
	return arr
}

func Cap(arr []string, i int) []string {
	if i-1 >= 0 && i+1 < len(arr) && string(arr[i]+arr[i+1]) == "(cap)" {
		for j := 0; j < i; j++ {
			if i-1-j >= 0 && IsLetter(arr[i-1-j]) {
				arr[i-1-j] = strings.ToLower(arr[i-1-j])
				arr[i-1-j] = strings.Title(arr[i-1-j])
				break
			}
		}
		arr[i], arr[i+1] = " ", " "
	} else if i+3 < len(arr) && string(arr[i]+arr[i+1]) == "(cap," {
		if IsNumber(arr[i+2]) {
			num := TypeChanger(arr[i+2])
			for j := 1; j <= num; j++ {
				if i >= num && i-j >= 0 {
					if IsLetter(arr[i-j]) {
						arr[i-j] = strings.ToLower(arr[i-j])
						arr[i-j] = strings.Title(arr[i-j])
					} else {
						num++
					}
				} else {
					flag = false
					return nil
				}
				arr[i], arr[i+1], arr[i+2], arr[i+3] = " ", " ", " ", " "
			}
		}
	}
	return arr
}

func Hex(arr []string, i int) []string {
	var num int64
	var err error
	ind := 0
	if i-1 >= 0 && string(arr[i]+arr[i+1]) == "(hex)" {
		for j := 0; j < i; j++ {
			if i-1-j >= 0 && IsLetter(arr[i-1-j]) || IsNumber(arr[i-1-j]) {
				num, err = strconv.ParseInt(arr[i-1-j], 16, 64)
				ind = i - 1 - j
				break
			}
		}
		if err != nil {
			flag = false
			return nil
		}
		arr[ind] = strconv.FormatInt(num, 10)
		arr[i], arr[i+1] = " ", " "
	}
	return arr
}

func Bin(arr []string, i int) []string {
	var num int64
	var err error
	ind := 0
	if i-1 >= 0 && string(arr[i]+arr[i+1]) == "(bin)" {
		for j := 0; j < i; j++ {
			if i-1-j >= 0 && IsNumber(arr[i-1-j]) {
				num, err = strconv.ParseInt(arr[i-1-j], 2, 64)
				ind = i - 1 - j
				break
			}
		}
		if err != nil {
			flag = false
			return nil
		}
		arr[ind] = strconv.FormatInt(num, 10)
		arr[i], arr[i+1] = " ", " "
	}
	return arr
}

func TypeChanger(str string) int {
	num, err := strconv.Atoi(str)
	if err == nil {
		return num
	}
	return -1
}

func IsLetter(s string) bool {
	flagletter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			flagletter = true
		}
	}
	return flagletter
}

func IsNumber(s string) bool {
	flagletter := false
	for _, r := range s {
		if unicode.IsNumber(r) {
			flagletter = true
		}
	}
	return flagletter
}

func IsPunct(r rune) bool {
	if unicode.IsPunct(r) {
		return true
	}
	return false
}

func Quotation(arr []string) []string {
	flagquot := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == "'" && !flagquot && i+1 < len(arr) && arr[i+1] != "'" {
			arr[i] = " "
			arr[i+1] = "'" + arr[i+1]
			flagquot = true

		} else if arr[i] == "'" && flagquot && i-1 >= 0 {
			flagquot = false
			arr[i] = "  "
			arr[i-1] = arr[i-1] + "'"
		} else if arr[i] == "'" && !flagquot && i+1 < len(arr) && arr[i+1] == "'" {
			flagquot = true
		}
	}

	return arr
}
