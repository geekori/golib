/*

The library is to simulate the try catch statement.

usage:

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

 */
package exception

import "reflect"

// exception struct，describe information related to exceptions
type Exception struct {
	Id int       // exception id
	Msg string   // exception msg
}
//
type TryStruct struct {
	// Simulate the catch chain part of the exception statement
	// key: the exception id
	// value: the function that handles the catch part
	catches map[int]ExceptionHandler
	// the function that handle the try part
	try   func()
}
// Simulate the try part of the exception statement
func Try(tryHandler func()) *TryStruct {
	tryStruct := TryStruct{
		catches: make(map[int]ExceptionHandler),
		try: tryHandler,
	}
	return &tryStruct
}

// the function type that handles the catch part
type ExceptionHandler func(Exception)



// the function that handles the catch part
// exceptionId: exception id
// catch: catch function
func (this *TryStruct) Catch(exceptionId int, catch func(Exception)) *TryStruct {
	this.catches[exceptionId] = catch
	return this
}
// the function that handles the finally part
func (this *TryStruct) Finally(finally func()) {
	defer func() {
		if e := recover(); nil != e {
			if reflect.TypeOf(e).String() == "Exception" {
				// The recover function returns the value of the Exception type.
				// This value is also the parameter value of the panic function
				exception := e.(Exception)
				// Search for a specific function that handles exceptions based on the exception id
				if catch, ok := this.catches[exception.Id]; ok {
					// Call a function that handles exceptions
					catch(exception)
				}
			} else {
				// The exception thrown by the system，exception id is -1
				exception := Exception{-1,e.(error).Error()}
				// if the exception handler is specified，it will be called.
				if catch, ok := this.catches[-1]; ok {
					catch(exception)
				}

			}

			// Call the function that handles the finally part
			finally()
		}
	}()


	// Call the function that handles the try part
	this.try()
}
// Function that throws an exception
// id: exception id
// msg: exception message
func Throw(id int, msg string) Exception {
	panic(Exception{id,msg})
}

