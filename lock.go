package main
import "sync"

var l sync.Mutex
var a string
var done bool

func routine() {

	a = "Hi from inside go routine"
        // l.Lock()
	done = true
        l.Unlock()

}

func main(){

        l.Lock()
	go routine()  // kick off green thread
        l.Lock()

	for !done {
	}

    println(a)
}

