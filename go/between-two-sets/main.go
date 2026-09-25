package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getTotalX(a []int32, b []int32) int32 {
	candidates:
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	panicOnError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMulitpleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMulitpleInput[0], 10, 64)
	panicOnError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMulitpleInput[1], 10, 64)
	panicOnError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	arr := extractIntArray(arrTemp, n)

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	brr := extractIntArray(brrTemp, m)

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)
	writer.Flush()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	line, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.Trim(string(line), "\r\n")
}

func extractIntArray(arr []string, size int32) []int32 {
	var result []int32
	for i := 0; i < int(size); i++ {
		item, err := strconv.ParseInt(arr[i], 10, 64)
		panicOnError(err)
		result = append(result, int32(item))
	}
	return result
}
