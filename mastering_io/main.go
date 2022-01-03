package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go_practice/mastering_io/imgcat"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "missing paths of imagecat")
		os.Exit(2)
	}

	if err := cat(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "could not cat %s: %v\n", os.Args[1], err)
	}

	//examples()
}

func cat(path string) error {
	imgFile, err := os.Open(path)
	if err != nil {
		return errors.New(fmt.Sprintf("could not open image, err : %v", err))
	}
	defer imgFile.Close()

	wc := imgcat.NewWriter(os.Stdout)
	if _, err = io.Copy(wc, imgFile); err != nil {
		return err
	}
	return wc.Close()
}

func examples() {
	fmt.Println("ioPipe Example:")
	ioPipeExample()
	fmt.Println("\nmultiReader Example:")
	multiReaderExample()
	fmt.Println("\nMultiWriter Example:")
	multiWriterExample()
}

func ioPipeExample() {
	//important: if you write something into pw, then it should be read from pr synchronously - vice versa
	pr, pw := io.Pipe()
	go func() {
		_, err := fmt.Fprint(pw, "I ma mastering io!")
		if err != nil {
			panic("Error on writing to io writer")
		}
		// closing the writer will send EOF to the reader synchronously.
		// `fatal error: all goroutines are asleep - deadlock!` error will be thrown if the writer is not closed
		// because reader will be indefinitely waiting for EOF event
		pw.Close()
	}()

	if _, err := io.Copy(os.Stdout, pr); err != nil {
		panic("Error on copying reader to stdout")
	}
}

func multiReaderExample() {
	header := strings.NewReader("<msg>")
	msg := strings.NewReader("Hello")
	footer := strings.NewReader("</msg>")

	for _, r := range []io.Reader{header, msg, footer} {
		_, err := io.Copy(os.Stdout, r)
		if err != nil {
			panic(err)
		}
	}

	//this is same as above
	//mr := io.MultiReader(header, msg, footer)
	//io.Copy(os.Stdout, mr)
}

func multiWriterExample() {
	buf := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, os.Stderr, buf)
	_, err := fmt.Fprint(mw, "Hello\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("buff : %v\n", buf)
}
