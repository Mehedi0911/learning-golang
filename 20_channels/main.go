package main

func main() {
	messageChan := make(chan string)

	messageChan <- "Hello, channels!"
	msg := <-messageChan
	println(msg)

}
