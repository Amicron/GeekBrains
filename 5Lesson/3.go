package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "fmt"
    "io"
)

/*Syntax Go. HomeWork 5
*Andrey Shuranov, dated May 12, 2019
*/

func main() {
    // Load a TXT file.
    f, _ := os.Open("C:\\Users\\limeh\\Desktop\\first\\5Lesson\\file.txt")

    // Create a new reader.
    r := csv.NewReader(bufio.NewReader(f))
    for {
        record, err := r.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }
        // Display record.
        // ... Display record length.
        // ... Display all individual elements of the slice.
        fmt.Println(record)
        fmt.Println(len(record))
        for value := range record {
            fmt.Printf("  %v\n", record[value])
        }
    }
}