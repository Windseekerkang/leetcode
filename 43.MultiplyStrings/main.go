package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(multiply("123", "456"))
}

//var biM  = map[byte]int{
//	'0':0,
//	'1':1,
//	'2':2,
//	'3':3,
//	'4':4,
//	'5':5,
//	'6':6,
//	'7':7,
//	'8':8,
//	'9':9,
//}

//func multiply(num1 string, num2 string) string {
//	result:=0
//	len1,len2:=len(num1),len(num2)
//	for i := len1 - 1; i >= 0; i-- {
//		for j := len2 - 1; j >= 0; j-- {
//			c1:=biM[num1[i]]
//			c2:=biM[num2[j]]
//			c3:=int(math.Pow10(len1+len2-i-j-2))
//			result = result + c1*c2*c3
//		}
//	}
//	if result == 0 {
//		return "0"
//	}
//	ibm:=make(map[int]byte,len(biM))
//	for key, value := range biM {
//		ibm[value] = key
//	}
//	bs:=make([]byte,0)
//	for result>0 {
//		bs = append(bs,ibm[result%10])
//		result = result/10
//	}
//	l,r:=0,len(bs)-1
//	for l < r {
//		bs[l],bs[r] = bs[r],bs[l]
//		l++
//		r--
//	}
//	return string(bs)
//}

func multiply(num1 string, num2 string) string {
	num1ints, num2ints, result := make([]int, len(num1)), make([]int, len(num2)), make([]int, len(num2)+len(num1))
	for key, v := range num1 {
		num1ints[key] = int(v - '0')
	}
	for key, v := range num2 {
		num2ints[key] = int(v - '0')
	}

	for k1, v1 := range num1ints {
		for k2, v2 := range num2ints {
			mul, p1, p2 := v1*v2, k1+k2, k1+k2+1
			sum := mul + result[p2]

			result[p1] += sum / 10
			result[p2] = sum % 10
		}
	}

	for i := len(result) - 1; i > 0; i-- {
		result[i-1] += result[i] / 10
		result[i] = result[i] % 10
	}

	s := "0123456789"
	buffer := bytes.NewBuffer([]byte{})
	for _, value := range result {
		if len(buffer.Bytes()) == 0 && value == 0 {
			continue
		}
		buffer.WriteByte(s[value])
	}
	if len(buffer.Bytes()) == 0 {
		return "0"
	}
	return buffer.String()
}
