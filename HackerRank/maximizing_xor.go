/*

Given two integers, l and r, find the maximal value of a xor b, where a and b satisfy the following condition

l <= a <= b <= r

*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the maximizingXor function below.
func maximizingXor(l int32, r int32) int32 {
    var max int32
    for a := l; a <= r; a++ {
        for b := a; b <= r; b++ {
            if a^b > max {
                max = a^b
            }
        }
    }
    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    lTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    l := int32(lTemp)

    rTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    r := int32(rTemp)

    result := maximizingXor(l, r)

    fmt.Fprintf(writer, "%d\n", result)

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
