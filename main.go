package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	nconstArg := os.Args[1]

	query := "SELECT primaryName FROM names WHERE nconst = @nconst LIMIT 1;"

	nconst := exec(query, nconstArg)

	fmt.Printf("nconst: %q\n", nconst)
}

func exec(query string, arg string) string {
	fname := "name.basics.tsv"

	f, err := os.Open(fname)

	if err != nil {
		fmt.Printf("ERR: %q", err)
		os.Exit(-1)
	}

	r := bufio.NewReader(f)

	//skip headers
	_, _, err = r.ReadLine()
	if err != nil {
		fmt.Printf("ERR: %q", err)
		os.Exit(-1)
	}

	for err != io.EOF {
		l, _, err := r.ReadLine()
		if err != nil && err != io.EOF {
			fmt.Printf("ERR: %q", err)
			os.Exit(-1)
		}
		if err == io.EOF {
			return "NOT FOUND"
		}
		fields := strings.Split(string(l), "\t")

		if fields[0] == arg {
			return fields[1]
		}
	}
	return "NOT_FOUND"
}
