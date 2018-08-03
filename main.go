package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {

	// must put the http://
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*
		bs := make([]byte, 9999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	//io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	t := triangle{
		base:   2,
		height: 5,
	}

	s := square{
		sideLength: 4,
	}

	printArea(t)
	printArea(s)

	//fmt.println()

	// resp sería lo que se lee del archivo
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}

func (t triangle) getArea() float64 {
	return 05 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
