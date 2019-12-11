package main

import (
    "fmt"
    "os"
)

func main() {

    var e int

    fmt.Scanf("%d", &e)
    fmt.Println(e)
    fmt.Fprintf(os.Stdout, "%#X\n", e)

}