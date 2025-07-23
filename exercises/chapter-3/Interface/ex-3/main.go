package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type Alterer interface {
	Notifier
	RaiseAlert(lever int, message string)
}

type SystemMonitor struct {}

func (m SystemMonitor) Notify(message string) {
	fmt.Println("Notification:", message)
}

func (a SystemMonitor) RaiseAlert(level int, message string) {
	fmt.Println("Alert message:", message)
	fmt.Println("Level of alert:", level)
}

func main() {
	sistMon:=SystemMonitor{}
	var notif Notifier = sistMon
	notif.Notify("hello helllo")
	var alert Alterer = sistMon
	alert.Notify("everything ok")
	alert.RaiseAlert(2,"not ok")

	var myInterface Alterer
	fmt.Println(myInterface)
	myInterface.Notify("test")
	// the runtime error ocures because the memory address is invalid or nil pointer 
	// and we have an nil interface so it does not know what type it is and it's triggering
	// the panic build-in function

}
