package main

import (
	"fmt"
	"time"

	"magic8"
)

// var failureMessage = "Can you ask me a question?"

func main() {

	// if len(os.Args) < 2 {
	// 	return
	// }
	fmt.Println(magic8.New(time.Now().Unix()).Answer())

}
