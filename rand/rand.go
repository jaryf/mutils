package rand

import (
	"encoding/binary"
	"time"
)

var (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52
	symbols    = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                   // 32
	digits     = "0123456789"                                           // 10
	characters = letters + digits + symbols                             // 94
)

// Intn 方法返回大于等0且不大于max的随机整数，即：[0, max)
func Intn(max int) int {
	if max <= 0 {
		return max
	}
	n := int(binary.LittleEndian.Uint32(<-bufferChan)) % max
	if (max > 0 && n < 0) || (max < 0 && n > 0) {
		return -n
	}
	return n
}

// B 方法用于返回指定长度的二进制[]byte数据。
func B(n int) []byte {
	if n <= 0 {
		return nil
	}
	i := 0
	b := make([]byte, n)
	for {
		copy(b[i:], <-bufferChan)
		i += 4
		if i >= n {
			break
		}
	}
	return b
}

// N 返回min到max之间的随机数
func N(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		return Intn(max-min+1) + min
	}
	if min < 0 {
		return Intn(max+(0-min)+1) - (0 - min)
	}
	return 0
}

// S 方法用于返回指定长度的数字、字符，第二个参数symbols用于指定知否返回的随机字符串中包含特殊字符。
func S(n int, symbols ...bool) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = B(n)
	)
	for i := range b {
		if len(symbols) > 0 && symbols[0] {
			b[i] = characters[numberBytes[i]%94]
		} else {
			b[i] = characters[numberBytes[i]%62]
		}
	}
	return string(b)
}

// D
func D(min, max time.Duration) time.Duration {
	multiple := int64(1)
	if min != 0 {
		for min%10 == 0 {
			multiple *= 10
			min /= 10
			max /= 10
		}
	}
	n := int64(N(int(min), int(max)))
	return time.Duration(n * multiple)
}

// Str 用于从给定的字符列表中选择指定长度的随机字符串返回，并且支持unicode字符
func Str(s string, n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b     = make([]rune, n)
		runes = []rune(s)
	)
	if len(runes) <= 255 {
		numberBytes := B(n)
		for i := range b {
			b[i] = runes[int(numberBytes[i])%len(runes)]
		}
	} else {
		for i := range b {
			b[i] = runes[Intn(len(runes))]
		}
	}
	return string(b)
}

// Digits 返回指定长度的随机数字字符串
func Digits(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = B(n)
	)
	for i := range b {
		b[i] = digits[numberBytes[i]%10]
	}
	return string(b)
}

// Letters 返回指定长度的随机英文字符串
func Letters(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = B(n)
	)
	for i := range b {
		b[i] = letters[numberBytes[i]%52]
	}
	return string(b)
}

// Symbols 返回指定长度的随机特殊字符串
func Symbols(n int) string {
	if n <= 0 {
		return ""
	}
	var (
		b           = make([]byte, n)
		numberBytes = B(n)
	)
	for i := range b {
		b[i] = symbols[numberBytes[i]%32]
	}
	return string(b)
}

// Perm returns, as a slice of n int numbers, a pseudo-random permutation of the integers [0,n).
func Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}

// Meet randomly calculate whether the given probability `num`/`total` is met.
func Meet(num, total int) bool {
	return Intn(total) < num
}

// MeetProb randomly calculate whether the given probability is met.
func MeetProb(prob float32) bool {
	return Intn(1e7) < int(prob*1e7)
}
