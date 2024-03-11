package interfaces

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Test interface {
	getTest() string
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type MyTest struct{}

// Implement the getTest method for MyTest
func (mt MyTest) getTest() string {
	return "test"
}
