package main

import (
    "fmt"
    "os"
    "strconv"
    "log"
)

func safeScanln(onError string) string {
    var answer string;
    _, err := fmt.Scanln(&answer)

    if err != nil {
        log.Print(err)
        return onError
    } else {
        return answer
    }
}

func readArg(value string) int {
    converted_integer, err := strconv.Atoi(value)
    if err != nil {
        fmt.Printf("Cannot convert '%s' into an integer, %s\n", value, err)
        fmt.Printf("Enter another value: ")

        return readArg(safeScanln("0"))
    } else {
        return converted_integer
    }
}

func main() {
    arguments := os.Args

    sum := 0

    for _, value := range arguments[1:len(arguments) - 1] {
        sum += readArg(value)
    }

    fmt.Println("Sum:", sum)
}
