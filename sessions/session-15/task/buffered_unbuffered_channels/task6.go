package buffered_unbuffered_channels

//Write a program that:
//Creates an unbuffered channel of type string.
//In the main function, send a value ("Hello") into the channel without launching a Goroutine to receive it.
//Observe that the program will block (deadlock) because there is no receiver.
//Expected Outcome: The program will block and fail with a runtime error:

func NoReverseChanel(ch chan string) {

}
