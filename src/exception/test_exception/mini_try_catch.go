package main

import (
	"log"
	"fmt"
	"reflect"
)
type Exception struct {
	Id int
	Msg string
}


func Try(tryHandler func()) *TryStruct {
	tryStruct := TryStruct{
		catches: make(map[int]ExceptionHandler),
		try: tryHandler,
	}
	return &tryStruct
}

type ExceptionHandler func(Exception)

type TryStruct struct {
	catches map[int]ExceptionHandler
	try   func()
}

func (this *TryStruct) Catch(exceptionId int, catch func(Exception)) *TryStruct {
	this.catches[exceptionId] = catch
	return this
}

func (this *TryStruct) Finally(finally func()) {
	defer func() {

		if e := recover(); nil != e {

			if reflect.TypeOf(e).String() == "Exception" {
				exception := e.(Exception)
				if catch, ok := this.catches[exception.Id]; ok {

					catch(exception)
				}
			} else {
				exception := Exception{-1,e.(error).Error()}

				if catch, ok := this.catches[-1]; ok {
					catch(exception)
				}

			}

			finally()
		}
	}()

	this.try()
}

func Throw(id int, msg string) Exception {
	panic(Exception{id,msg})
}

func main() {

	Try(func() {
		//log.Println("try...")
		var i = 0
		fmt.Println(4/i)
		//Throw(2,"error2")
	}).Catch(1, func(e Exception) {
		log.Println(e.Id,e.Msg)
	}).Catch(2, func(e Exception) {
		log.Println(e.Id,e.Msg)
	}).Catch(-1, func(e Exception) {
		log.Println(e.Id,e.Msg)
	}).Finally(func() {
		log.Println("finally")
	})

}
