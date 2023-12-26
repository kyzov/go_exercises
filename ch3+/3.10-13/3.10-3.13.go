package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s = "1230000"
	fmt.Println(comma(s)) //3.10

	var s1 = "1230000.1251515"
	fmt.Println(comma2(s1)) //3.11

	fmt.Println(angram())

}

func comma(s string) string { //3.10
	var n bytes.Buffer
	var a = 0
	for i := len(s) - 1; i >= 0; i-- {

		if a == 3 {
			n.WriteString(",")
			a = 0
		}
		a++
		n.WriteString(string(s[i]))
	}
	b := n.Bytes()
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func comma2(s string) string { //3.11
	//var n bytes.Buffer
	var t = ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) != "." {
			t += string(s[i])
		} else {
			break
		}
	}
	return comma(t) + s[len(t):]
}

func angram() string { // 3.12
	s1 := "12345"
	s2 := "54321"
	for i := 0; i < len(s1); i++ {
		if string(s1[i]) != string(s2[len(s2)-i-1]) {
			return "no"
		}
	}
	return "yes"
}

const ( //3.13
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)
