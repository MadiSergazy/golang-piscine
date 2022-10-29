package main

func ReverseBits(oct byte) byte {
	var result byte = 0
	for onybyte := 8; onybyte > 0; onybyte-- {
		result = (result << 1) | (oct & 1)
		oct >>= 1
	}
	return result
}
