package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 121
	fmt.Println(isPalindrome(a))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n := strconv.Itoa(x)
	fmt.Println(n)
	if n[0] == n[1] {
		return true
	}
	return false
}

// func main() {
// 	nums1 := []int{1, 2}
// 	nums2 := []int{3, 4}
// 	fmt.Println(findMedianSortedArrays(nums1, nums2))
// }

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	for _, v := range nums2 {
// 		nums1 = append(nums1, v)
// 	}
// 	for i := 0; i < len(nums1); i++ {
// 		for j := 0; j < len(nums1); j++ {
// 			if nums1[i] < nums1[j] {
// 				nums1[i], nums1[j] = nums1[j], nums1[i]
// 			}
// 		}
// 	}
// 	var n float64
// 	if len(nums1)%2 == 1 {
// 		n = float64(nums1[(len(nums1) / 2)])
// 	} else if len(nums1)%2 == 0 {
// 		n = float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1]) / 2
// 	}
// 	return float64(n)
// }

// func Atoi(s string) int {
// 	n := 0
// 	var negative bool
// 	if s[0] == '-' {
// 		s = s[1:]
// 		negative = true
// 	}
// 	for _, w := range s {
// 		if negative {
// 			n = n*10 + (int(w-48) * -1)
// 		} else {
// 			n = n*10 + int(w-48)
// 		}
// 	}
// 	return n
// }

// func Itoa(n int) string {
// 	res := ""
// 	for n != 0 {
// 		res = string(n%10+48) + res
// 		n = n / 10
// 	}

// 	return res
// }

// func main() {
// 	arg := os.Args[1:]
// 	nbr := 0
// 	for _, ch := range arg {
// 		for _, smb := range ch {
// 			nbr = nbr*10 + int(rune(smb-'0'))
// 		}
// 	}
// }

//////////////////////////*  PrintHex   *///////////////////////////////

// func main() {
// 	s := "5156454"
// 	fmt.Println(printhex(s))
// }

// func Atoi(s string) int {
// 	res := 0

// 	for _, val := range s {
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res
// }

// func printhex(s string) string {
// 	base := "0123456789abcdef"
// 	lenBase := len(base)
// 	num := Atoi(s)
// 	res := ""

// 	for num > 0 {
// 		res = string(base[num%lenBase]) + res
// 		num /= lenBase
// 	}

// 	return res
// }

///////////////////////*  wdmatch  *////////////////////////////////

// func main() {
// 	if len(os.Args) == 3 {
// 		str1 := os.Args[1]
// 		str2 := os.Args[2]
// 		i := 0
// 		word := ""
// 		for _, val := range str2 {
// 			if val == rune(str1[i]) {
// 				word += string(val)
// 				i++
// 			}
// 			if len(word) == len(str1) {
// 				break
// 			}
// 		}

// 		if word == os.Args[1] {
// 			for _, val := range word {
// 				z01.PrintRune(val)
// 			}
// 		} else {
// 			return
// 		}
// 	}

// 	z01.PrintRune('\n')
// }

//////////////////////////////* Chunk *//////////////////////////////

// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7, 12, 2, 23, 14}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }

// func Chunk(slice []int, size int) {
// 	if size != 0 {
// 		res := make(map[int][]int, 8)
// 		i := 0
// 		for _, val := range slice {
// 			res[i] = append(res[i], val)
// 			if len(res[i]) == size {
// 				i++
// 			}
// 		}
// 		fmt.Print("[")
// 		for i := 0; i < len(res); i++ {
// 			if i < len(res)-1 {
// 				fmt.Printf("%v ", res[i])
// 			} else {
// 				fmt.Printf("%v", res[i])
// 			}
// 		}
// 		fmt.Print("]")
// 		fmt.Println()
// 	} else {
// 		fmt.Println()
// 	}
// }

////////////////////////////* Searchreplace *///////////////////////////

// func main() {
// 	if len(os.Args) == 4 {
// 		s := os.Args[1]
// 		first := os.Args[2]
// 		sec := os.Args[3]
// 		res := ""

// 		for _, val := range s {
// 			if val == rune(first[0]) {
// 				res += string(sec[0])
// 			} else {
// 				res += string(val)
// 			}
// 		}
// 		fmt.Println(res)
// 	}
// }

///////////////////////////* tabmult *////////////////////////////

// func main() {
// 	if len(os.Args) > 1 {
// 		num := Atoi(os.Args[1])

// 		for i := 1; i < 10; i++ {
// 			res := i * num
// 			fmt.Printf("%d x %d = %d\n", i, num, res)
// 		}
// 	}
// }

// func Atoi(s string) (res int) {
// 	for _, val := range s {
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res
// }

/////////////////////////* Roman Numbers *///////////////////////////

// type roman struct {
// 	num        int
// 	romanDigit string
// }

// func main() {
// 	if len(os.Args) > 1 {
// 		str := os.Args[1]
// 		number, err := Atoi(str)
// 		if err != nil || number > 3999 || number == 0 {
// 			fmt.Println("ERROR: cannot convert to roman digit")
// 			return
// 		}

// 		romanDigits := []roman{
// 			{num: 1000, romanDigit: "M"},
// 			{num: 900, romanDigit: "CM"},
// 			{num: 500, romanDigit: "D"},
// 			{num: 400, romanDigit: "CD"},
// 			{num: 100, romanDigit: "C"},
// 			{num: 90, romanDigit: "XC"},
// 			{num: 50, romanDigit: "L"},
// 			{num: 40, romanDigit: "XL"},
// 			{num: 10, romanDigit: "X"},
// 			{num: 9, romanDigit: "IX"},
// 			{num: 5, romanDigit: "V"},
// 			{num: 4, romanDigit: "IV"},
// 			{num: 1, romanDigit: "I"},
// 		}

// 		result, romanSeq := MakeRomanNumber(number, romanDigits)
// 		slice := Split(romanSeq, "+")
// 		divide := DoubleRoman(slice)

// 		fmt.Printf("%v\n%v\n", divide, result)
// 	} else {
// 		fmt.Println("Not enough argument")
// 	}
// }

// func MakeRomanNumber(number int, romanDigits []roman) (string, string) {
// 	result, romanSeq := "", ""
// 	for _, val := range romanDigits {
// 		for number >= val.num {
// 			result += val.romanDigit
// 			romanSeq += val.romanDigit + "+"
// 			number = number - val.num
// 		}
// 	}
// 	return result, romanSeq[:len(romanSeq)-1]
// }

// func DoubleRoman(slice []string) string {
// 	res := ""
// 	for _, val := range slice {
// 		if len(val) == 2 {
// 			res += fmt.Sprintf("(%s-%s)+", string(val[1]), string(val[0]))
// 		} else {
// 			res += val + "+"
// 		}
// 	}

// 	return res[:len(res)-1]
// }

// func Split(romanSeq, plus string) []string {
// 	var res []string
// 	word := ""
// 	for _, val := range romanSeq {
// 		if string(val) == plus {
// 			res = append(res, word)
// 			word = ""
// 		} else {
// 			word += string(val)
// 		}
// 	}
// 	res = append(res, word)
// 	return res
// }

// func Atoi(s string) (res int, err error) {
// 	for _, val := range s {
// 		if val < 48 || val > 57 {
// 			return 0, err
// 		}
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res, nil
// }

///////////////////////

// func main() {
// 	var (
// 		i     int
// 		nums  []int
// 		nums2 []int
// 	)
// 	for len(nums) < 5 {
// 		fmt.Scan(&i)
// 		nums = append(nums, i)
// 	}
// 	fmt.Println("1 slice: ", nums)
// 	nums2 = append(nums2, nums[2:]...)
// 	fmt.Println("2 slice: ", nums2)
// }

/////////////////////* Alpha Mirror *///////////////////////

// func main() {
// 	if len(os.Args) > 1 {
// 		str := os.Args[1]
// 		res := ""

// 		for _, val := range str {
// 			if val >= 'A' && val <= 'Z' {
// 				res += string('Z' - val + 'A')
// 			} else if val >= 'a' && val <= 'z' {
// 				res += string('z' - val + 'a')
// 			} else {
// 				res += string(val)
// 			}
// 		}

// 		for _, val := range res {
// 			z01.PrintRune(val)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

/////////////////////* Comapre *////////////////////////

// func main() {
// 	fmt.Println(Compare("Hello!", "Hoooo!"))
// 	fmt.Println(Compare("Salut!", "lut!"))
// 	fmt.Println(Compare("Ola!", "Ol"))
// }

// func Compare(s1, s2 string) (res int) {
// 	if s1 > s2 {
// 		return 1
// 	} else if s1 < s2 {
// 		return -1
// 	}
// 	return 0
// }

////////////////////////* Doop *//////////////////////////////

// const (
// 	MaxUint = ^uint(0)
// 	MaxInt  = int(MaxUint >> 1)
// )

// func main() {
// 	if len(os.Args) == 4 {
// 		num1, err := Atoi(os.Args[1])
// 		if err != nil {
// 			return
// 		}
// 		num2, err2 := Atoi(os.Args[3])
// 		if err2 != nil {
// 			return
// 		}
// 		sign := os.Args[2]
// 		Doop(num1, num2, sign)
// 	}
// }

// func Doop(num1, num2 int, sign string) {
// 	switch {
// 	case sign == "+":
// 		if num1 > 0 && num2 > 0 && num1+num2 < 0 {
// 			return
// 		}
// 		fmt.Println(num1 + num2)
// 	case sign == "-":
// 		if num1 < num2 && num1-num2 > 0 {
// 			return
// 		}
// 		fmt.Println(num1 - num2)
// 	case sign == "/":
// 		if num2 == 0 {
// 			fmt.Println("No division by 0")
// 		} else {
// 			fmt.Println(num1 / num2)
// 		}
// 	case sign == "%":
// 		if num2 == 0 {
// 			fmt.Println("No modulo by 0")
// 		} else {
// 			fmt.Println(num1 % num2)
// 		}
// 	case sign == "*":
// 		if num1*num2 > MaxInt/2 || num1 > 0 && num2 > 0 && num1*num2 < 0 {
// 			return
// 		}
// 		fmt.Println(num1 * num2)
// 	}
// }

// func Atoi(s string) (res int, err error) {
// 	flag := 1
// 	for _, val := range s {
// 		if val == '-' {
// 			flag = -1
// 		} else if val < 48 || val > 57 {
// 			return 0, errors.New("error")
// 		} else {
// 			res = res*10 + (int(val) - 48)
// 		}
// 	}

// 	if res < 0 {
// 		return 0, errors.New("error")
// 	}

// 	return res * flag, nil
// }

///////////////////////* FindPrevPrime *///////////////////////////

// func main() {
// 	fmt.Println(FindPrevPrime(100))
// 	fmt.Println(FindPrevPrime(2))
// }

// func FindPrevPrime(num int) int {
// 	for num > 1 {
// 		if isPrime(num) {
// 			return num
// 		}
// 		num--
// 	}
// 	return 0
// }

// func isPrime(num int) bool {
// 	d := 2
// 	if num <= 1 {
// 		return false
// 	}
// 	for d*d <= num {
// 		if num%d == 0 {
// 			return false
// 		}
// 		d += 1
// 	}
// 	return true
// }

//////////////////////* is Power of 2 *///////////////////////

// func isPowerOf2(num int) bool {
// 	return num != 0 && (num&(num-1) == 0)
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		num := Atoi(os.Args[1])
// 		fmt.Println(isPowerOf2(num))
// 	}
// }

// func Atoi(s string) (res int) {
// 	for _, val := range s {
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res
// }

/////////////////////////* Union */////////////////////////

// func main() {
// 	if len(os.Args) == 3 {
// 		str1 := OnlyUnique(os.Args[1])
// 		str2 := os.Args[2]

// 		for _, val := range str2 {
// 			if notThere(str1, val) {
// 				str1 += string(val)
// 			}
// 		}
// 		fmt.Println(str1)
// 	} else {
// 		fmt.Println()
// 	}
// }

// func notThere(s string, let rune) bool {
// 	for _, val := range s {
// 		if val == let {
// 			return false
// 		}
// 	}
// 	return true
// }

// func OnlyUnique(s string) (res string) {
// 	for _, val := range s {
// 		if notThere(res, val) {
// 			res += string(val)
// 		}
// 	}
// 	return res
// }

////////////////////////* inter *///////////////////////

// func main() {
// 	if len(os.Args) == 3 {
// 		str1 := OnlyUnique(os.Args[1])
// 		str2 := os.Args[2]
// 		res := ""
// 		for _, val := range str1 {
// 			for _, val2 := range str2 {
// 				if val == val2 {
// 					res += string(val2)
// 					break
// 				}
// 			}
// 		}
// 		fmt.Println(res)
// 	} else {
// 		fmt.Println()
// 	}
// }

// func OnlyUnique(s string) (res string) {
// 	for _, val := range s {
// 		if notThere(res, val) {
// 			res += string(val)
// 		}
// 	}
// 	return res
// }

// func notThere(s string, let rune) bool {
// 	for _, val := range s {
// 		if val == let {
// 			return false
// 		}
// 	}
// 	return true
// }

///////////////////////* PrintBits *//////////////////////

// func main() {
// 	if len(os.Args) == 2 {
// 		slice := make([]int, 8)
// 		num, err := Atoi(os.Args[1])
// 		if err != nil {
// 			for _, val := range slice {
// 				z01.PrintRune(rune(val) + 48)
// 			}
// 			z01.PrintRune('\n')
// 			return
// 		}
// 		i := len(slice) - 1
// 		for num != 0 {
// 			slice[i] = num % 2
// 			num /= 2
// 			i--
// 		}

// 		for _, val := range slice {
// 			z01.PrintRune(rune(val) + 48)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func Atoi(s string) (res int, err error) {
// 	for _, val := range s {
// 		if val < 48 || val > 57 {
// 			return 0, errors.New("error")
// 		}
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res, nil
// }

////////////////////////* Swap Bits *///////////////////////

// func SwapBits(octet byte) byte {
// 	return octet<<4 + octet>>4
// }

//////////////////////* Reverse Bits *///////////////////////

// package main

// func ReverseBits(oct byte) byte {
// 	var result byte = 0
// 	for onybyte := 8; onybyte > 0; onybyte-- {
// 		result = (result << 1) | (oct & 1)
// 		oct >>= 1
// 	}
// 	return result
// }

///////////////////////////* Pig latin *//////////////////////

// func main() {
// 	if len(os.Args) == 2 {
// 		vowels := "aeiouAEIOU"
// 		str := os.Args[1]
// 		res := ""
// 		for i, val := range str {
// 			if isVowel(val, vowels) {
// 				res += str[i:] + str[:i]
// 			}
// 		}
// 		if res == "" {
// 			noVowels(res)
// 			return
// 		}
// 		res += "ay"
// 		for _, val := range res {
// 			z01.PrintRune(val)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func noVowels(s string) {
// 	str := "No vowels"
// 	for _, val := range str {
// 		z01.PrintRune(val)
// 	}
// 	z01.PrintRune('\n')
// }

// func isVowel(let rune, vowels string) bool {
// 	for _, val := range vowels {
// 		if val == let {
// 			return true
// 		}
// 	}
// 	return false
// }

/////////////////////* Repeat Alpha *//////////////////////////

// func main() {
// 	if len(os.Args) == 2 {
// 		str := os.Args[1]
// 		i := 0
// 		for _, val := range str {
// 			if val >= 'a' && val <= 'z' {
// 				i = int(val - 96)
// 			} else if val >= 'A' && val <= 'Z' {
// 				i = int(val - 64)
// 			} else {
// 				z01.PrintRune(val)
// 				continue
// 			}
// 			for j := 0; j < i; j++ {
// 				z01.PrintRune(val)
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

/////////////////////////* SortWordArr *//////////////////////////

// func main() {
// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	SortWordArr(result)
// }

// func SortWordArr(a []string) {
// 	for i := 0; i < len(a)-1; i++ {
// 		for j := 0; j < len(a)-i-1; j++ {
// 			if a[j] > a[j+1] {
// 				a[j], a[j+1] = a[j+1], a[j]
// 			}
// 		}
// 	}
// }

//////////////////////////* PrintHex */////////////////////////////

// func main() {
// 	if len(os.Args) == 2 {
// 		num, err := Atoi(os.Args[1])
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println(PrintHex(num))
// 	}
// }

// func PrintHex(num int) string {
// 	hex := "0123456789abcdef"
// 	length := len(hex)
// 	res := ""

// 	for num > 0 {
// 		res = string(hex[num%length]) + res
// 		num /= length
// 	}
// 	return res
// }

// func Atoi(s string) (res int, err error) {
// 	for _, val := range s {
// 		if val < '0' || val > '9' {
// 			return 0, errors.New("ERROR")
// 		}
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res, nil
// }

//////////////////////////* GCD *///////////////////////////

// func main() {
// 	if len(os.Args) == 3 {
// 		num1 := Atoi("42")
// 		num2 := Atoi("10")

// 		gcd := GCD(num1, num2)
// 		res := Itoa(gcd)

// 		for _, val := range res {
// 			z01.PrintRune(val)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func GCD(num1, num2 int) int {
// 	for num2 != 0 {
// 		num1 = num1 % num2
// 		num1, num2 = num2, num1
// 	}
// 	return num1
// }

// func Atoi(s string) (res int) {
// 	for _, val := range s {
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res
// }

// func Itoa(num int) string {
// 	var slice []int
// 	res := ""

// 	for num != 0 {
// 		slice = append(slice, num%10)
// 		num /= 10
// 	}

// 	for i := len(slice) - 1; i >= 0; i-- {
// 		res += string(rune(slice[i]) + 48)
// 	}

// 	return res
// }

////////////////////////* Hiddenp */////////////////////////

// func main() {
// 	if len(os.Args) == 3 {
// 		str1 := os.Args[1]
// 		str2 := os.Args[2]
// 		i := 0
// 		res := ""
// 		for _, val := range str2 {
// 			if len(res) == len(str1) {
// 				break
// 			}
// 			if val == rune(str1[i]) {
// 				res += string(val)
// 				i++
// 			}
// 		}
// 		if res == str1 {
// 			fmt.Println(1)
// 		} else {
// 			fmt.Println(0)
// 		}
// 	}
// }

///////////////////////////* Rostring *///////////////////////////

// func main() {
// 	if len(os.Args) == 2 {
// 		str := Split(os.Args[1])

// 		for _, val := range str[1:] {
// 			Print(val)
// 			z01.PrintRune(' ')
// 		}
// 		Print(str[0])
// 		z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune('\n')
// 	}
// }

// func Print(s string) {
// 	for _, val := range s {
// 		z01.PrintRune(val)
// 	}
// }

// func Split(str string) []string {
// 	var res []string
// 	word := ""

// 	for _, val := range str {
// 		if val == ' ' && word != "" {
// 			res = append(res, word)
// 			word = ""
// 		} else if val != ' ' {
// 			word += string(val)
// 		}
// 	}
// 	if word != "" {
// 		res = append(res, word)
// 	}
// 	return res
// }

///////////////////////* Split */////////////////////

// func main() {
// 	s := "HelloBONKhowBONKareBONKyou?"
// 	fmt.Printf("%#v\n", Split(s, "BONK"))
// }

// func Split(s, sep string) []string {
// 	var res []string
// 	word := ""
// 	i := 0

// 	for i < len(s) {
// 		if s[i] == sep[0] {
// 			frag := s[i : i+len(sep)]
// 			if frag == sep {
// 				res = append(res, word)
// 				word = ""
// 				i += len(sep)
// 				continue
// 			}
// 		}
// 		word += string(s[i])
// 		i++
// 	}
// 	if word != "" {
// 		res = append(res, word)
// 	}
// 	return res
// }

//////////////////////////* RevwStr */////////////////////////

// func main() {
// 	if len(os.Args) == 2 {
// 		str := Split(os.Args[1])

// 		for i := len(str) - 1; i >= 0; i-- {
// 			if i != 0 {
// 				fmt.Printf("%v ", str[i])
// 			} else {
// 				fmt.Printf("%v", str[i])
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func Split(str string) []string {
// 	var res []string
// 	word := ""

// 	for _, val := range str {
// 		if val == ' ' && word != "" {
// 			res = append(res, word)
// 			word = ""
// 		} else if val != ' ' {
// 			word += string(val)
// 		}
// 	}
// 	if word != "" {
// 		res = append(res, word)
// 	}
// 	return res
// }

////////////////////////////////* Fprime *//////////////////////////////

// func main() {
// 	if len(os.Args) == 2 {
// 		num, err := Atoi(os.Args[1])
// 		if err != nil || num == 0 || num == 1 {
// 			return
// 		}
// 		seq := []int{}

// 		for i := 2; i <= num; i++ {
// 			for num%i == 0 {
// 				seq = append(seq, i)
// 				num /= i
// 			}
// 		}
// 		s := ""
// 		for i, val := range seq {
// 			if i != len(seq)-1 {
// 				s += Itoa(val) + "*"
// 			} else {
// 				s += Itoa(val)
// 			}
// 		}
// 		for _, val := range s {
// 			z01.PrintRune(val)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func Itoa(num int) string {
// 	var slice []int
// 	res := ""

// 	for num != 0 {
// 		slice = append(slice, num%10)
// 		num /= 10
// 	}

// 	for i := len(slice) - 1; i >= 0; i-- {
// 		res += string(rune(slice[i]) + 48)
// 	}

// 	return res
// }

// func Atoi(s string) (res int, err error) {
// 	for _, val := range s {
// 		if val < '0' || val > '9' {
// 			return 0, errors.New("ERROR")
// 		}
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res, nil
// }

//////////////////////////////* AddPrimeSum *//////////////////////////////

// func main() {
// 	if len(os.Args) == 2 {
// 		num, err := Atoi(os.Args[1])
// 		if err != nil || num == 0 {
// 			z01.PrintRune('0')
// 		}
// 		res := 0
// 		for i := 2; i <= num; i++ {
// 			if isPrime(i) {
// 				res += i
// 			}
// 		}
// 		strNum := Itoa(res)
// 		for _, val := range strNum {
// 			z01.PrintRune(val)
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// }

// func isPrime(num int) bool {
// 	d := 2
// 	if num <= 1 {
// 		return false
// 	}
// 	for d*d <= num {
// 		if num%d == 0 {
// 			return false
// 		}
// 		d += 1
// 	}
// 	return true
// }

// func Atoi(s string) (res int, err error) {
// 	for _, val := range s {
// 		if val < '0' || val > '9' || val == '-' {
// 			return 0, errors.New("ERROR")
// 		}
// 		res = res*10 + (int(val) - 48)
// 	}
// 	return res, nil
// }

// func Itoa(num int) string {
// 	var slice []int
// 	res := ""

// 	for num != 0 {
// 		slice = append(slice, num%10)
// 		num /= 10
// 	}

// 	for i := len(slice) - 1; i >= 0; i-- {
// 		res += string(rune(slice[i]) + 48)
// 	}

// 	return res
// }

////////////////////////////* AtoiBase *//////////////////////////

// func main() {
// 	fmt.Println(AtoiBase("125", "0123456789"))
// 	fmt.Println(AtoiBase("1111101", "01"))
// 	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
// 	fmt.Println(AtoiBase("uoi", "choumi"))
// 	fmt.Println(AtoiBase("bbbbbab", "-ab"))
// }

// func AtoiBase(s string, base string) int {
// 	if base[0] == '-' {
// 		return 0
// 	}
// 	lenBase := len(base)
// 	lenNum := len(s)
// 	res := 0

// 	for i := 0; i < len(s); i++ {
// 		res += Index(rune(s[i]), base) * RecursivePower(lenBase, lenNum-1)
// 		lenNum--
// 	}
// 	return res
// }

// func Index(r rune, s string) int {
// 	for i, val := range s {
// 		if val == r {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func RecursivePower(lenBase, lenNum int) int {
// 	if lenNum < 0 {
// 		return 0
// 	}

// 	if lenNum == 0 {
// 		return 1
// 	}
// 	return lenBase * RecursivePower(lenBase, lenNum-1)
// }

/////////////////////////////* PrintMemory *///////////////////////////////

// func main() {
// 	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
// }

// func PrintMemory(arr [10]byte) {
// 	str := ""
// 	for i, val := range arr {
// 		if val == 0 {
// 			fmt.Printf("%d%d", 0, 0)
// 		} else {
// 			fmt.Printf("%s", printhex(int(val)))
// 		}

// 		if (i+1)%4 == 0 || i == len(arr)-1 {
// 			fmt.Println()
// 		} else {
// 			fmt.Print(" ")
// 		}

// 		if val >= '!' && val <= '~' {
// 			str += string(rune(val))
// 		} else {
// 			str += "."
// 		}
// 	}

// 	fmt.Println(str)
// }

// func printhex(num int) string {
// 	base := "0123456789abcdef"
// 	lenBase := len(base)
// 	res := ""

// 	for num > 0 {
// 		res = string(base[num%lenBase]) + res
// 		num /= lenBase
// 	}
// 	return res
// }

//////////////////////////////////* RPNcalc *////////////////////////////////
