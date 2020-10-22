package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func testDemo2() {
	input := "foo\nbar\nbaz"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Not actually needed since it’s a default split function.
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func testScanner() {
	input := "foo  bar    baz"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}


func main() {

	f, err := os.Open("/Users/yuqiang/go_workspace/go_learning/src/gopl.io/ch7/bytecounter/main.go")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	info, err := f.Stat()

	var data []byte = make([]byte, info.Size())
	input := "foo   bar      baz"
	f.Read(data)
	adv, token, err := bufio.ScanWords([]byte(input), false)
	fmt.Println(adv)
	fmt.Println(string(token))

	//testScanner()
	testDemo2()

}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(strings.NewReader(string(p)))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		*c += 1
	}
	return len(p), nil
}
