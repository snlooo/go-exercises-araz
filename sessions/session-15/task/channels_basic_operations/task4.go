package channels_basic_operations

//Task 4: Channel Range Iteration
//
//Write a program that:
//Creates an unbuffered channel of type int.
//Launches a Goroutine that sends values from 1 to 5 into the channel and closes it.
//The main function should receive and print the values from the channel using a for loop.

func SendValueAndClose(ch chan int) chan int {

	for i := 1; i <= 5; i++ {
		go func() {
			ch <- i
		}()
	}
	return ch
}
