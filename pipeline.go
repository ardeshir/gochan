// Stages using multiple channels
package main
import("fmt")


func main() {
	// set up the pipeline and consume the output
	for transformedData := range processStage(inputStage(2,3)) {
		fmt.Println(transformedData)
	}
}

// Stage One
func inputStage(data ...int) <-chan int {
	outPutData := make(chan int)
	go func() {
	     for _, inputData := range data {
		   outPutData <- inputData

             }
		close(outPutData)
	}()
	return outPutData
}


// Stage Two
func processStage(input <-chan int) <-chan int {
	outPutData := make(chan int)
	go func() {
	   for inputData := range input {
             outPutData <- inputData * inputData
	   }
             close(outPutData)
	}()
	return outPutData
}

