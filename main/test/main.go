package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 뒷쪽 아이와 비교하여 내 캔디개수 변경
func adjustCandy(i int, arr []int32, candyCount []int) {
	if i < 0 { return }

	// 내 랭킹이 뒷쪽아이보다 높은데, 캔디가 같거나 작다면 내꺼 하나 증가시키고 다시 앞쪽 비교
	if (arr[i] > arr[i+1]) && (candyCount[i] <= candyCount[i+1]) {
		candyCount[i]++
		adjustCandy(i-1, arr, candyCount)
	}

	return
}

// Complete the candies function below.
func candies(n int32, arr []int32) int64 {
	// 캔디개수 배열
	candyCount := make([]int, n)

	// 맨 앞 사람에게 1개 줌
	candyCount[0] = 1

	// 내 캔디의 개수를 정함
	for i := 1; i < int(n); i++ {
		// 앞에 사람보다 내 랭킹이 높으면 캔디 + 1, 그 밖에는 1개만 주고, 앞쪽 아이들 조정
		candyCount[i] = 1
		if arr[i] > arr[i-1] {
			candyCount[i] = candyCount[i-1] + 1
		} else {

			adjustCandy(i-1, arr, candyCount)
		}
	}
	//for i := 1; i < int(n); i++ {
	//	if candyCount[i] = 1; arr[i] > arr[i-1] {
	//		candyCount[i] = candyCount[i-1] + 1
	//	}
	//
	//}

	// 모든 캔디 합계
	var sum int
	fmt.Println(candyCount)
	for i := 0; i < len(candyCount); i++ {
		sum = sum + int(candyCount[i])
	}

	return int64(sum)
}

func main() {
	dat, err := os.Open("./input.txt")
	checkError(err)

	reader := bufio.NewReaderSize(dat, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	n := int32(nTemp)
	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		//fmt.Println(arrItemTemp)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := candies(n, arr)

	fmt.Println(result)
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
