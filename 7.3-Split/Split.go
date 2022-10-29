package main

import "fmt"

func main() {
	s := "HelloBONKhowBONKareBONKyou?"
	fmt.Printf("%#v\n", Split(s, "BONK"))
}

func Split(s, sep string) []string {
	var res []string
	word := ""
	i := 0

	for i < len(s) {
		if s[i] == sep[0] {
			frag := s[i : i+len(sep)]
			if frag == sep {
				res = append(res, word)
				word = ""
				i += len(sep)
				continue
			}
		}
		word += string(s[i])
		i++
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}
