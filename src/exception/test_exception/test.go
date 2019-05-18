package main

import (
	"exception"
	"log"
)

func main() {

	exception.Try(func() {
		log.Println("try...")
		exception.Throw(2,"error2")
	}).Catch(1, func(e exception.Exception) {
		log.Println(e.Id,e.Msg)
	}).Catch(2, func(e exception.Exception) {
		log.Println(e.Id,e.Msg)
	}).Finally(func() {
		log.Println("finally")
	})
}

