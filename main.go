package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	fmt.Println("Input some strings:")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	a := scan.Text()
	if a == reverseString(a) {
		fmt.Printf("%s - palindrome\n", a)
	} else {
		fmt.Printf("%s - not a palindrome\n", a)
	}
}

func reverseString(s string) string {
	runa := []rune(s)
	leng := len(runa)
	for i := 0; i < leng/2; i++ {
		j := leng - i - 1
		runa[i], runa[j] = runa[j], runa[i]
	}
	return string(runa)
}   