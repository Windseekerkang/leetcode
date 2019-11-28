package main

import (
	"bytes"
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println(intToRoman5(1995))
}

func intToRoman1(num int) string {
	var buffer bytes.Buffer
	for i := 3; i >= 0; i-- {
		m := math.Pow10(i)
		n := math.Floor(float64(num) / m)
		num = int(math.Mod(float64(num), m))
		if n != 0 {
			getRomanString(int(n), int(m), &buffer)
		}
	}
	return buffer.String()
}

func getRomanString(n int, m int, buffer *bytes.Buffer) {
	switch m {
	case 1000:
		for i := 0; i < n; i++ {
			buffer.WriteString("M")
		}
	case 100:
		switch n {
		case 9:
			buffer.WriteString("CM")
		case 8:
			buffer.WriteString("DCCC")
		case 7:
			buffer.WriteString("DCC")
		case 6:
			buffer.WriteString("DC")
		case 5:
			buffer.WriteString("D")
		case 4:
			buffer.WriteString("CD")
		case 3:
			buffer.WriteString("CCC")
		case 2:
			buffer.WriteString("CC")
		default:
			buffer.WriteString("C")
		}
	case 10:
		switch n {
		case 9:
			buffer.WriteString("XC")
		case 8:
			buffer.WriteString("LXXX")
		case 7:
			buffer.WriteString("LXX")
		case 6:
			buffer.WriteString("LX")
		case 5:
			buffer.WriteString("L")
		case 4:
			buffer.WriteString("XL")
		case 3:
			buffer.WriteString("XXX")
		case 2:
			buffer.WriteString("XX")
		default:
			buffer.WriteString("X")
		}
	default:
		switch n {
		case 9:
			buffer.WriteString("IX")
		case 8:
			buffer.WriteString("VIII")
		case 7:
			buffer.WriteString("VII")
		case 6:
			buffer.WriteString("VI")
		case 5:
			buffer.WriteString("V")
		case 4:
			buffer.WriteString("IV")
		case 3:
			buffer.WriteString("III")
		case 2:
			buffer.WriteString("II")
		default:
			buffer.WriteString("I")
		}
	}

}

func intToRoman5(num int) string {
	buffer := make([]byte, 0, 15)
	for i := 3; i >= 0; i-- {
		m := math.Pow10(i)
		n := int(math.Floor(float64(num) / m))
		num = int(math.Mod(float64(num), m))
		if n != 0 {
			switch m {
			case 1000:
				for i := 0; i < n; i++ {
					buffer = append(buffer, 'M')
				}
			case 100:
				switch n {
				case 9:
					buffer = append(buffer, 'C', 'M')
				case 8:
					buffer = append(buffer, 'D', 'C', 'C', 'C')
				case 7:
					buffer = append(buffer, 'D', 'C', 'C')
				case 6:
					buffer = append(buffer, 'D', 'C')
				case 5:
					buffer = append(buffer, 'D')
				case 4:
					buffer = append(buffer, 'C', 'D')
				case 3:
					buffer = append(buffer, 'C', 'C', 'C')
				case 2:
					buffer = append(buffer, 'C', 'C')
				default:
					buffer = append(buffer, 'C')
				}
			case 10:
				switch n {
				case 9:
					buffer = append(buffer, 'X', 'C')
				case 8:
					buffer = append(buffer, 'L', 'X', 'X', 'X')
				case 7:
					buffer = append(buffer, 'L', 'X', 'X')
				case 6:
					buffer = append(buffer, 'L', 'X')
				case 5:
					buffer = append(buffer, 'L')
				case 4:
					buffer = append(buffer, 'X', 'L')
				case 3:
					buffer = append(buffer, 'X', 'X', 'X')
				case 2:
					buffer = append(buffer, 'X', 'X')
				default:
					buffer = append(buffer, 'X')
				}
			default:
				switch n {
				case 9:
					buffer = append(buffer, 'I', 'X')
				case 8:
					buffer = append(buffer, 'V', 'I', 'I', 'I')
				case 7:
					buffer = append(buffer, 'V', 'I', 'I')
				case 6:
					buffer = append(buffer, 'V', 'I')
				case 5:
					buffer = append(buffer, 'V')
				case 4:
					buffer = append(buffer, 'I', 'V')
				case 3:
					buffer = append(buffer, 'I', 'I', 'I')
				case 2:
					buffer = append(buffer, 'I', 'I')
				default:
					buffer = append(buffer, 'I')
				}
			}
		}
	}
	return *(*string)(unsafe.Pointer(&buffer))
}

func intToRoman3(num int) string {
	romans := make([]byte, 0, 15)
	i := 3
	for num != 0 {
		m := int(math.Pow10(i))
		n := num / m
		num = num % m

		switch n {
		case 9:
			romans = append(romans, romanNumerals[i][0], romanNumerals[i+1][0])
		case 4:
			romans = append(romans, romanNumerals[i][0], romanNumerals[i][1])
		default:
			if n >= 5 {
				romans = append(romans, romanNumerals[i][1])
				n -= 5
			}
			for n > 0 {
				romans = append(romans, romanNumerals[i][0])
				n -= 1
			}
		}
		i--
	}
	return *(*string)(unsafe.Pointer(&romans))
}

var romanNumerals [4][2]byte = [4][2]byte{
	{'I', 'V'},
	{'X', 'L'},
	{'C', 'D'},
	{'M', ' '},
}
