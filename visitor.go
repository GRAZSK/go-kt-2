package main

import "fmt"

// Элементы
type Shape interface {
	Accept(v Visitor)
}

type Circle struct{ Radius float64 }
func (c *Circle) Accept(v Visitor) { v.VisitCircle(c) }

type Square struct{ Side float64 }
func (s *Square) Accept(v Visitor) { v.VisitSquare(s) }

// Посетитель
type Visitor interface {
	VisitCircle(c *Circle)
	VisitSquare(s *Square)
}

// Конкретный посетитель 1: Вычисление площади
type AreaCalculator struct { TotalArea float64 }

func (a *AreaCalculator) VisitCircle(c *Circle) {
	a.TotalArea += 3.14 * c.Radius * c.Radius
}
func (a *AreaCalculator) VisitSquare(s *Square) {
	a.TotalArea += s.Side * s.Side
}

// Конкретный посетитель 2: Экспорт в строку
type XMLExporter struct { XML string }

func (x *XMLExporter) VisitCircle(c *Circle) {
	x.XML += fmt.Sprintf("<circle radius='%f'/>", c.Radius)
}
func (x *XMLExporter) VisitSquare(s *Square) {
	x.XML += fmt.Sprintf("<square side='%f'/>", s.Side)
}

func main() {
	shapes := []Shape{&Circle{Radius: 5}, &Square{Side: 4}}

	areaCalc := &AreaCalculator{}
	for _, s := range shapes { s.Accept(areaCalc) }
	fmt.Println("Total Area:", areaCalc.TotalArea)

	xmlExp := &XMLExporter{}
	for _, s := range shapes { s.Accept(xmlExp) }
	fmt.Println("XML:", xmlExp.XML)
}