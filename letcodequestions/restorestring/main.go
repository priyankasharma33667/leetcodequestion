package main

import "fmt"

func restoreString(s string, indices []int) string {
	n := len(s)
	fmt.Println("*********", n)
	restored := make([]byte, n)
	fmt.Println("*********restore", restored)

	for i := 0; i < n; i++ {
		restored[indices[i]] = s[i]
	}

	return string(restored)

}

func main() {
	s := "priyanka"
	indices := []int{0, 1, 2, 7, 4, 3, 5, 6}
	fmt.Println(restoreString(s, indices))

}
