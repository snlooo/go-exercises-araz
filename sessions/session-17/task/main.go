package main

import (
	"session-17/ctx_pkg"
	"session-17/rate_limit_and_throttling"
)

func main() {

	//Task 1
	ctx_pkg.PrintNumbers()

	//Task 2
	ctx_pkg.SimulateCancellation()

	//Task 3
	rate_limit_and_throttling.PrintMessage()

	//Task 4
	rate_limit_and_throttling.JobStatus()
}
