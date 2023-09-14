/*
	Посетитель.
	Вид: Структурный.
	Суть паттерна - добавление нового фунцкионала в программу, не изменяя классы, надо которыми
	эти операции совершаются.
	+: упрощает добавление операций, работающих со сложными структурами объектов
	+: объединяет родственные операции в одном классе
	-: паттерн не оправдан, если иерархия элементов часто меняется
	-: может привести к нарушению инкапсуляции элементов
*/
package main

import "fmt"

// Element - интерфейс элемента, который может быть посещен посетителем
type Element interface {
	accept(visitor Visitor)
}

// ConcreteElementA - конкретный элемент A
type ConcreteElementA struct {
	name string
}

func NewConcreteElementA(name string) *ConcreteElementA {
	return &ConcreteElementA{
		name: name,
	}
}

func (e *ConcreteElementA) accept(visitor Visitor) {
	visitor.visitConcreteElementA(e)
}

// ConcreteElementB - конкретный элемент B
type ConcreteElementB struct {
	value int
}

func NewConcreteElementB(value int) *ConcreteElementB {
	return &ConcreteElementB{
		value: value,
	}
}

func (e *ConcreteElementB) accept(visitor Visitor) {
	visitor.visitConcreteElementB(e)
}

// Visitor - интерфейс посетителя
type Visitor interface {
	visitConcreteElementA(element *ConcreteElementA)
	visitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor - конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) visitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("Visited ConcreteElementA with name %s\n", element.name)
}

func (v *ConcreteVisitor) visitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("Visited ConcreteElementB with value %d\n", element.value)
}

// ObjectStructure - объектная структура, которая содержит элементы и использует посетителя для их обхода
type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) addElement(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) removeElement(element Element) {
	for i, e := range os.elements {
		if e == element {
			os.elements = append(os.elements[:i], os.elements[i+1:]...)
			break
		}
	}
}

func (os *ObjectStructure) accept(visitor Visitor) {
	for _, element := range os.elements {
		element.accept(visitor)
	}
}

func main() {
	elementA := NewConcreteElementA("elementA")
	elementB := NewConcreteElementB(42)

	objectStructure := &ObjectStructure{}
	objectStructure.addElement(elementA)
	objectStructure.addElement(elementB)

	visitor := &ConcreteVisitor{}
	objectStructure.accept(visitor)
}
