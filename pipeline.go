/*  Pipeline using multiple channels   ******************/
/*
 1) We can create pipelines by chaining channels
 2) A chained channel is a stage in the pipeline
 3) A stage will receive from an inbound channel until closed
 4) Pipelines are a design pattern not an explicit type

*********************************************************/

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

