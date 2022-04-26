package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func process(file string, counts map[string]int) error {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	entries := strings.Fields(string(contents))
	for _, e := range entries {
		if e != "" {
			if _, ok := counts[e]; !ok {
				counts[e] = 0
			}
			counts[e] = counts[e] + 1
		}
	}
	return nil
}

func processAll(files []string, counts map[string]int) (err error) {
	for _, arg := range files {
		go func() {
			err = process(arg, counts)
		}()
	}
	return nil
}

func main() {

	array1 := [][]int{{1, 2, 3, 4}, {4, 4, 5}}

	fmt.Println("caP : ", cap(array1))
	fmt.Println("Len : ", len(array1))

	counts := make(map[string]int, 0)

	err := processAll(os.Args[1:], counts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", counts)
}
