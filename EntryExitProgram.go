package main

import "fmt"

var activeUserCount int = 0

func EntryExit() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}

func entry() {
	activeUserCount++
}

func exit() {
	activeUserCount--
}
