//Реализация паттерна фасад
//	Вид: Структурный.
//	Суть паттерна - предоставление простого или урезанного интерфейса для работы
//	со сложной подсистемой (фреймворка например).
//	+: изоляция клиента от сложной подсистемы, тем самым ее облегченное использование
//	-: риск создания перегруженного объекта
package main

import "fmt"

type CPU struct{}

func (c *CPU) start() {
	fmt.Println("Starting CPU")
}

func (c *CPU) execute() {
	fmt.Println("Executing CPU instructions")
}

func (c *CPU) stop() {
	fmt.Println("Stopping CPU")
}

type Ram struct{}

func (m *Ram) load() {
	fmt.Println("Loading data into ram")
}

type HardDrive struct{}

func (hd *HardDrive) read() {
	fmt.Println("Reading data from hard drive")
}

type ComputerFacade struct {
	cpu       *CPU
	ram       *Ram
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		ram:       &Ram{},
		hardDrive: &HardDrive{},
	}
}

func (cf *ComputerFacade) start() {
	cf.cpu.start()
	cf.ram.load()
	cf.hardDrive.read()
	cf.cpu.execute()
}

func (cf *ComputerFacade) stop() {
	cf.cpu.stop()
}

func main() {
	computer := NewComputerFacade()
	computer.start()
	computer.stop()
}
