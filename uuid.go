package main

import (
    "crypto/rand"
    "fmt"
)

func random_filename_16_char() (s string, err error) {
    b := make([]byte, 8)
    _, err = rand.Read(b)
    if err != nil {
        return
    }
    s = fmt.Sprintf("%x", b)
    return
}

func main() {
    s, _ := random_filename_16_char()
    fmt.Println(s)
}
