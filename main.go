package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("no input is given")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("can't open the file")
		return
	}

	I := 0
	N, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("number is expected")
		return
	}
	for {
		buf := make([]byte, N)
		n, err := f.Read(buf)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("can't read from the file")
			return
		}

		fmt.Printf("%08x: ", I*N)
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
