package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var CandyAmount []int
var total int

// 뒷쪽 아이와 비교하여 내 캔디개수 변경
func adjustCandy(i int, arr []int32, candyCount []int) {
	if i < 0 { return }
	// 내 랭킹이 뒷쪽아이보다 높은데, 캔디가 같거나 작다면 내꺼 하나 증가시키고 다시 앞쪽 비교
	if (arr[i] > arr[i+1]) && (candyCount[i] <= candyCount[i+1]) { candyCount[i]++;	adjustCandy(i-1, arr, candyCount) }
	return
}
// Complete the candies function below.
// n is 아이수
// arr is 여러 아이의 score array
func candies(n int32, arr []int32) int64 {

	for i := 1; i < int(n); i++ { CandyAmount[i] = 1; if adjustCandy(i-1, arr, CandyAmount); arr[i] > arr[i-1] { CandyAmount[i] = CandyAmount[i-1] + 1 } }

	return int64(sum(CandyAmount))
}

func sum(score []int) int {
	for i := 0; i < len(score); i++ { total += score[i] }
	return total
}

func main() {

	// 아이들의 score 값 읽어오기
	dat, err := os.Open("./test_writer.txt")
	checkError(err)

	// 파일 읽기
	reader := bufio.NewReaderSize(dat, 1024*1024)

	// 맨처음 숫자를 아이들 수라고 생각함
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	var arr []int32

	// 그 이후 숫자들을 아이들 점수라고 생각하고 배열에 넣음
	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	// 캔디개수 배열
	CandyAmount = make([]int, n)

	// 맨 앞 사람에게 1개 줌
	CandyAmount[0] = 1

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
