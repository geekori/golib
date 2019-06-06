package main

import (
	"exception"
	"log"
	"fmt"
)

func main() {

	exception.Try(func() {
		log.Println("try...")
		//exception.Throw(2,"error2")
		var i = 0
		fmt.Println(10/i)
	}).Catch(1, func(e exception.Exception) {
		log.Println(e.Id,e.Msg)
	}).Catch(2, func(e exception.Exception) {
		log.Println(e.Id,e.Msg)
	}).Catch(-1, func(e exception.Exception) {
		log.Println(e.Id,e.Msg)
	}).Finally(func() {
		log.Println("finally")
	})
}

