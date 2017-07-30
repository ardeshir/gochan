package main
import (
     "fmt"
    // "time"
)

var unbuff = make(chan int)
var   buff = make(chan int, 5)

func main() {
    count :=5
    go func() {
	for i := 0; i <= count; i++ {
		fmt.Println("Sending message: ", i)
		buff <- i
	}
       close(buff)
    }()

    // time.Sleep(time.Second * 3)
    for num := range buff {
      fmt.Print("Receiving: ")
		fmt.Println(num)
    }
}
