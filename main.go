package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file := "./examples/long-nl.txt"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("can't open the file")
		return
	}
	var I = 0
	for {
		buf := make([]byte, 16)
		n, err := f.Read(buf)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("can't read from the file")
			return
		}

		fmt.Printf("%08x: ", I*16)
		for i := 0; i < n; i++ {
			if buf[i] < 0x10 {
				fmt.Print("0")
			}
			fmt.Print(strconv.FormatInt(int64(buf[i]), 16), " ")
		}
		I++

		// iterate on the buf and print each char, if the char is \n, \r or \t print . insteade
		for _, c := range buf {
			if c == '\n' || c == '\r' || c == '\t' || c == ' ' || c == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(c))
			}

		}
		fmt.Print("\n")
	}
}
