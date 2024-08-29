package library

import (
	"fmt"
	"math"
)

type Human interface {
	Display()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Display() {
	fmt.Println("Name:", p.Name, "Age:", p.Age)
}

func InterfaceFunction() {
	var h Human
	p := Person{
		Name: "John",
		Age:  25,
	}
	h = p
	h.Display()
	fmt.Println(h)

	var h1 Human = Person{Name: "John", Age: 25}
	h1.Display()

	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Println("Area of circle: ", getArea(circle))
	fmt.Println("Area of Rectanlge: ", getArea(rectangle))

	nodepod := NodePods{name: "node", version: "v10", npm: "npm"}
	pythonpod := PythonPods{name: "python", version: "v3.8", pip: "pip"}
	fmt.Println(getPods(nodepod))
	fmt.Println(getPods(pythonpod))
	getAllPods(nodepod, pythonpod)
}

// // ########################## Shape ####################################

type Shape interface {
	area() float64
}

/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a rectangle */
type Rectangle struct {
	width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

// ########################## Pods ####################################

type Pods interface {
	getPods() string
}

type NodePods struct {
	name    string
	version string
	npm     string
}

type PythonPods struct {
	name    string
	version string
	pip     string
}

func (np NodePods) getPods() string {
	return "NodePods " + np.name + " " + np.version + " " + np.npm
}

func (pp PythonPods) getPods() string {
	return "PythonPods " + pp.name + " " + pp.version + " " + pp.pip
}

func getPods(pod Pods) string {
	return pod.getPods()
}

func getAllPods(pods ...Pods) {
	for _, pod := range pods {
		fmt.Println(pod.getPods())
	}
}
