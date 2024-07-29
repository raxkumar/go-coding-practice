package main

import "fmt"

type EmployeeSalary interface {
	CalculateEmployeeSalary() float32
}

type Hr struct {
	baseSalary float32
	hrHike     float32
}

type Developer struct {
	baseSalary float32
	devHike    float32
	clientHike float32
}

func (hr Hr) CalculateEmployeeSalary() float32 {
	return hr.baseSalary * hr.hrHike * 2.50
}

func (dev Developer) CalculateEmployeeSalary() float32 {
	return dev.devHike * dev.clientHike * dev.baseSalary * 4.50
}

func PrintEmployeeSalary(es EmployeeSalary) {
	fmt.Println(es.CalculateEmployeeSalary())
}

func TestPolyorphism() {

	hr := Hr{120, 100}
	dev := Developer{150, 10, 50}

	PrintEmployeeSalary(hr)
	PrintEmployeeSalary(dev)

}
