package main

func main() {
	nums := []int{1, 2, 2, 4}
	findErrorNums(nums)
	/*
		args := os.Args[1:]
		if len(args) != 1 {
			return
		}
		san, err := strconv.Atoi(args[0])
		if err != nil || san == 2 {
			return
		}
		for i := 2; i < san; i++ {
			if san%i == 0 {
				fmt.Print(i, "*")
				san /= i
				i = 2
			}
		}
		fmt.Print(san, "\n")
	*/
}

func findErrorNums(nums []int) []int {

	mp := make(map[int]int)

	for _, value := range nums {
		mp[value] += 1
	}

	n := len(nums)
	var ans []int

	for i := 1; i <= n; i++ {
		freq, present := mp[i]

		if !present {
			ans = append(ans, i)
		}

		if present && freq == 2 {
			//prepend in array
			ans = append([]int{i}, ans...)
		}
	}
	return ans
}

/*{
	ar := os.Args
	if len(ar) != 2 {
		os.Exit(0)
	}
	n, er := strconv.Atoi(ar[1])
	if er != nil {
		os.Exit(0)
	}
	if n < 2 {
		return
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Print(i, "*")
			n /= i
			i = 1
		}
	}
	fmt.Print(n)
	fmt.Println()
}
*/
// func fprime(n int) {
// 	for x := 2; x < n*2; x++ {
// 		if n%x == 0 {
// 			fmt.Print(x)
// 			if n/x != 1 {
// 				fmt.Print("*")
// 			}
// 			fprime(n / x)
// 			break
// 		}
// 	}
// }

// func itoa(n int) string {
// 	sign := 1
// 	if n < 0 {
// 		n *= -1
// 		sign = -1
// 	}
// 	if n == 0 {
// 		return "0"
// 	}
// 	res := ""
// 	for n != 0 {
// 		res = string(rune(n%10+'0')) + res
// 		n /= 10
// 	}
// 	if sign == -1 {
// 		return "-" + res
// 	}
// 	return res
// }
