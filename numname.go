package main

import (
"fmt"
"bufio"
)

func main() {
    
    var age int
    fmt.Println("What is your age?")
    _, err: fmt.Scan(&age)

    
    reader := bufio.newReader(os.Stdin)
    var name string
    fmt.Println("What is your name?")
    name, _ := reader.readString("\n")

    fmt.Println("Your name is ", name, " and you are age ", age)
}
