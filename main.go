// Select on multiple channels
package main
import ("fmt"; "time")
func main() {
	c1 := getChannel("chan 1")
	c2 := getChannel("chan 2")
   for i := 1; i <= 4; i++ {
	select {
		case msg := <- c1:
                fmt.Println(msg)
		case mgs := <- c2:
		fmt.Println(mgs)
		}
 }
}

func getChannel(msg string) <- chan string {
	c := make(chan string)
	go func() {
	  for i := 1; i <= 2; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		// wait before sending next msg
		time.Sleep(150)
	   }
	}()
   return c 
}
