package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func hasDup(s []byte) bool {
	for i := 0; i < len(s); i++ {
		rc := s[i]
		if i == len(s)-1 {
			break
		}
		for _, rd := range s[i+1:] {
			if rd == rc {
				return true
			}
		}
	}
	return false
}

func main() {
	in, _ := ioutil.ReadAll(os.Stdin)
	var marker, marker2 int
	for i := 0; i < len(in); i++ {
		dup := hasDup(in[i : i+4])
		dup2 := hasDup(in[i : i+14])
		if !dup && marker == 0 {
			marker = i + 4
		}
		if !dup2 && marker2 == 0 {
			marker2 = i + 14
		}
		if marker != 0 && marker2 != 0 {
			break
		}
	}

	fmt.Println(marker, marker2)
}
