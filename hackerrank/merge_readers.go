package main

import (
	"bufio"
	// "bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
 * Complete the 'MergeReaders' function below.
 *
 * The function is expected to return an io.Reader and an error.
 * The function accepts following parameters:
 *  1. io.Reader r1
 *  2. io.Reader r2
 */

func MergeReaders(r1, r2 io.Reader) (io.Reader, error) {
	var result string

	input1 := make([]byte, 1)
	input2 := make([]byte, 1)

	for {
		// Read one character from the first reader
		n, err := r1.Read(input1)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if n > 0 {
			result += string(input1)
		}

		// Read one character from the second reader
		n, err = r2.Read(input2)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if n > 0 {
			result += string(input2)
		}
	}

	return strings.NewReader(result), nil

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	str1 := readLine(reader)

	str2 := readLine(reader)

	if len(str1) > len(str2) {
		str1 = str1[0:len(str2)]
	} else if len(str2) > len(str1) {
		str2 = str2[0:len(str1)]
	}

	r1 := strings.NewReader(str1)
	r2 := strings.NewReader(str2)

	resultReader, err := MergeReaders(r1, r2)
	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadAll(resultReader)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(writer, "%s\n", string(result))

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
