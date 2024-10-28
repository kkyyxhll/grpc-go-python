// // client.go
// package main

// func main() {
// 	addr := "192.168.2.30:22222"
// 	clientTest(addr)
// }

package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("FATAL")
		}
	}()

	panic("panic")
}
