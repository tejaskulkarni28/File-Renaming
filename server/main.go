package main

import (
	"fmt"
	"os"
)

var globalPath = ""
var newGlobalFileNamePath = ""

func getPath() {
	fmt.Println("Enter the path name you want to change: ")
	fmt.Scanln(&globalPath)

	fmt.Println("Enter the new path you want to change: ")
	fmt.Scanln(&newGlobalFileNamePath)

}

func changeFileName() {
	os.Rename(globalPath, newGlobalFileNamePath)
}

func main() {
	getPath()
	changeFileName()
}
