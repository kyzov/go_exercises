package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 6)
	fmt.Println(s)

	f := []string{"A", "A", "B", "B", "A", "C", "F", "F"}
	fmt.Println(removeDup(f))

	t := "1  +  1     =  2"
	fmt.Println(string(removeDupSpace([]byte(t))))

	k := "test tset"
	fmt.Println(string(reverseUTF8([]byte(k))))
}

func reverse(s *[5]int) { //4.3
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) { //4.4
	num := n % len(s)
	double := append(s, s[:num]...)
	copy(s, double[num:num+len(s)])
}

func removeDup(s []string) []string { //4.5
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}

func removeDupSpace(b []byte) []byte { //4.6
	var buf bytes.Buffer
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRuneInString(string(b[i:]))
		if unicode.IsSpace(r) {
			nextrune, _ := utf8.DecodeRuneInString(string(b[i+size:]))
			if !unicode.IsSpace(nextrune) {
				buf.WriteRune(' ')
			}
		} else {
			buf.WriteRune(r)
		}
		i += size
	}
	return buf.Bytes()
}

func reverseUTF8(b []byte) []byte { //4.7.1
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse2(b[i : i+size])
		i += size
	}
	reverse2(b)
	return b
}

func reverse2(b []byte) []byte { //4.7.2
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
