package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reader size setting
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	//os.Setenv("OUTPUT_PATH", "input")
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	// n is 아이 수
	n := int32(nTemp)

	var arr []int32

	for i := 0; i < int(n); i++ {

		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		// arrItem is 한 아이의 score
		arrItem := int32(arrItemTemp)

		// 여러 아이의 score array
		arr = append(arr, arrItem)

	}

	//fmt.Println(n)
	//fmt.Println(arr)
	fmt.Println(writer, "result")
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
