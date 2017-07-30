package main
import (
     "fmt"
     "time"
)

var unbuff = make(chan int)
var   buff = make(chan int)

func main() {
    count :=5
    go func() {
	for i := 0; i <= count; i++ {
		fmt.Println("Sending message: ")
		unbuff <- i
	}
    }()

    time.Sleep(time.Second * 3)
    for i := 0; i <= count; i++ {
      fmt.Println("Receiving: ")
		fmt.Println( <- unbuff)
    }
}
