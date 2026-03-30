package main

import "session-16/task/sync_atomic"

func main() {
	////Task 1
	//select_statement.PrintValuesFromChannels()
	////Task 2
	//ch := make(chan string, 1)
	////ch <- "hello"
	//select_statement.WaitWithTimeOut(ch)
	//
	////Task 3
	//sync_waitgroup.InformStarting()
	//
	////Task 4
	//ch4 := make(chan []int, 2)
	//var number = []int{1, 2, 3, 4, 5, 6, 7, 8}
	//ch4 <- number
	//sync_waitgroup.CalculateSum(ch4)

	//Task 5
	//sync_mutex.IncrementValue()

	//Task 6
	//ch := make(chan map[string]int, 1)
	//
	//students := map[string]int{
	//	"Garay":  32,
	//	"Ali":    34,
	//	"Medina": 23,
	//}
	//ch <- students
	//sync_mutex.UpdateGradeOfStudents(ch)

	//Task 7
	//ch := make(chan map[string]int, 1)
	//users := map[string]int{
	//	"Garay":  20,
	//	"Ali":    25,
	//	"Medina": 28,
	//}
	//
	//ch <- users
	//
	//sync_rwmutex.UpdateAgeOfUsers(ch)

	//Task 8
	//sync_rwmutex.IncrementValue(5)

	//Task 9
	sync_atomic.IncrementWithAtomic(10)
}
