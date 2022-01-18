package main

import (
	"fmt"

	"github.com/arturoeanton/gocommons/observer"
)

type ElementObservable struct {
	observer.BasicObservable
}

type ElementObserver struct {
	observer.BasicObserver
}

func (e *ElementObserver) Notify(data interface{}) {
	fmt.Println("State1:", data.(string))
}

func main() {
	observable1 := ElementObservable{}

	observer1 := ElementObserver{}
	observer1.GetID()

	observable1.AddObserver(&observer1)

	observable1.ChangeState("hola")

	observable1.ChangeState("hola1")

	fmt.Println(observable1.GetState())

}
