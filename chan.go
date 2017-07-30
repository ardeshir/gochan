package main

import (
    "fmt"
)

func main() {
   cWriter := make(chan int) // create a channel with make keyward

   // use a function literal as a goroutine channel writer
   go func() {
     cWriter <- 2
   } ()

   // use the main function goroutine as a channel reader
   cReader := <- cWriter  // read
   fmt.Println(cReader)
}


