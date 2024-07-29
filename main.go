package main

import "fmt"

type Device struct {
	name string
}

type Smartphone struct {
	Device
	version string
}

func (d Device) ErrorSound() {
	fmt.Println("NO DEAFULT SOUND")

}

// func (sm Smartphone) ErrorSound() {
// 	fmt.Println("BEEP BEEP")
// }

func main() {

	sm := Smartphone{
		Device:  Device{name: "apple"},
		version: "1.2.0",
	}

	sm.ErrorSound()
	sm.Device.ErrorSound()

}
