package main

import (
	"fmt"
	"time"
)

// func processNum(numChan chan int)  {
// 	for num := range numChan {
// 		fmt.Println("Processing number",num)
// 		time.Sleep(time.Second)
// 	}

// }

// func sum(result chan int, num1 int,  num2 int)  {
// 	newresult := num1 + num2
// 	result <- newresult

// }

// goroutine synchronizer
// func task(done chan bool)  {
// 	defer func () {done <- true}()
// 	fmt.Println("Processing...")
// }


func emailSender(emailChan chan string, done chan bool){
	defer func(){done <-true}()

	for email := range emailChan {
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}
}



func main() {
	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan,done)

	for i:= 0; i <= 5; i++{
		emailChan <- fmt.Sprintf("%d@gmail.com",i)
	}

	fmt.Println("done sending.")
	
	//this is done 
	close(emailChan)
	<-done

	emailChan <- "1@example"
	emailChan <- "2@example"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)
	<-done

	// done := make(chan bool) //unbuffer channel

	// go task(done)
	// <- done //block
	// result := make(chan int)

	// go sum(result,4,5)

	// res := <- result //blocking

	// fmt.Println(res)
	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	


	// messageChan := make(chan string)
	// // to Send the data to inside the channel
	// messageChan <- "ping"

	// // to Receive the data to inside the channel
	// msg := <- messageChan

	// fmt.Println(msg)
}