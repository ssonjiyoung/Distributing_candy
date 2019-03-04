package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var score []int32

// Complete the candies function below.
// n is 아이수
// arr is 여러 아이의 score array
func candies(n int32, arr []int32) int64 {

	// 여기에 구현하세요.
	// distributing candy
	fmt.Println(len(score))
	for i := 0; i < int(n); i++ {
		if arr[i] > arr[i+1] {
			score[i]++
		} else if arr[i] < arr[i+1] {

			score[i+1]++
		} else {
			if i == 0 {
				score[i]++
			} else {
				for j := i; j >= 0; j-- {
					score[i] = score[i-1]
					score[i]++
				}

			}
		}
	}
	result := total(score)
	return result
}
func total(score []int32) int64 {
	var total int32 = 0
	for i := 0; i < len(score); i++ {
		total += score[i]
	}
	return int64(total)

}
func main() {

	// reader size setting
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	os.Setenv("OUTPUT_PATH", "input")
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

	result := candies(n, arr)

	//fmt.Fprintf(writer, "%d\n", result)
	fmt.Println(writer, result)
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
