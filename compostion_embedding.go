package main

import "fmt"

type Machine struct {
	Name string
}

func (m *Machine) Sound() {
	fmt.Println("beep.. beep...")
}

type SmartPhone struct {
	machine Machine
	compnay string
}

func (m *SmartPhone) Sound() {
	fmt.Println("kek.. kek...")

}

func CompostionEmbedding() {

	sm := &SmartPhone{
		machine: Machine{"14-pro"},
		compnay: "apple",
	}
	// child sound
	sm.Sound()
	// parent sound
	sm.machine.Sound()

}
