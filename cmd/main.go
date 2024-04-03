package main

import (
	"eShop/db"
	"eShop/pkg/controller"
	"fmt"
	"os"
	
	"strings"
)

func main() {
	db.StartDbConnection()

	StartProgram()
}

func StartProgram() {
	args := os.Args[1:] // Получаем аргументы командной строки, начиная со второго элемента (индекс 1)
	argString := strings.Join(args, ", ")

	err := controller.BulkGetOrdersByIDs(argString)
	if err != nil {
		fmt.Println("Occurred error during getting information. Error is: ", err.Error())
	}
}
